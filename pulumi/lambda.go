package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createLambdas(ctx *pulumi.Context, usersTable *dynamodb.Table, assetsTable *dynamodb.Table) (*pulumi.StringOutput, *pulumi.String, error) {
	account, err := aws.GetCallerIdentity(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	region, err := aws.GetRegion(ctx, &aws.GetRegionArgs{})
	if err != nil {
		return nil, nil, err
	}

	// Create an IAM role.
	role, err := iam.NewRole(ctx, "task-exec-role", &iam.RoleArgs{
		AssumeRolePolicy: pulumi.String(`{
				"Version": "2012-10-17",
				"Statement": [{
					"Sid": "",
					"Effect": "Allow",
					"Principal": {
						"Service": "lambda.amazonaws.com"
					},
					"Action": "sts:AssumeRole"
				}]
			}`),
	})
	if err != nil {
		return nil, nil, err
	}

	// Attach a policy to allow writing logs to CloudWatch
	logPolicy, err := iam.NewRolePolicy(ctx, "lambda-log-policy", &iam.RolePolicyArgs{
		Role: role.Name,
		Policy: pulumi.String(`{
                "Version": "2012-10-17",
                "Statement": [{
                    "Effect": "Allow",
                    "Action": [
                        "logs:CreateLogGroup",
                        "logs:CreateLogStream",
                        "logs:PutLogEvents"
                    ],
                    "Resource": "arn:aws:logs:*:*:*"
                }]
            }`),
	})

	dynamoPolicy, err := iam.NewRolePolicy(ctx, "lambda-dynamo-policy", &iam.RolePolicyArgs{
		Role: role.Name,
		Policy: pulumi.Sprintf(`{
			"Version": "2012-10-17",
			"Statement": [
				{
					"Effect": "Allow",
					"Action": [
						"dynamodb:BatchGetItem",
						"dynamodb:BatchWriteItem",
						"dynamodb:DeleteItem",
						"dynamodb:GetItem",
						"dynamodb:PutItem"
					],
					"Resource": [
						"%s",	
						"%s"	
					]
				}
			]
		}`, assetsTable.Arn, usersTable.Arn),
	})

	// Set arguments for constructing the function resource.
	args := &lambda.FunctionArgs{
		Handler: pulumi.String("handler"),
		Role:    role.Arn,
		Runtime: pulumi.String("provided.al2023"),
		Code:    pulumi.NewFileArchive("../bin/lambda/api/api.zip"),
		// Arm64
		Architectures: pulumi.StringArray{
			pulumi.String("arm64"),
		},
	}

	// Create the lambda using the args.
	function, err := lambda.NewFunction(
		ctx,
		"basicLambda",
		args,
		pulumi.DependsOn([]pulumi.Resource{logPolicy, dynamoPolicy}),
	)
	if err != nil {
		return nil, nil, err
	}

	// Create a new API Gateway.
	gateway, err := apigateway.NewRestApi(ctx, "UpperCaseGateway", &apigateway.RestApiArgs{
		Name:        pulumi.String("UpperCaseGateway"),
		Description: pulumi.String("An API Gateway for the UpperCase function"),
		Policy: pulumi.String(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    },
    {
      "Action": "execute-api:Invoke",
      "Resource": "*",
      "Principal": "*",
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}`)})
	if err != nil {
		return nil, nil, err
	}

	// Add a resource to the API Gateway.
	// This makes the API Gateway accept requests on "/{message}".
	apiresource, err := apigateway.NewResource(ctx, "UpperAPI", &apigateway.ResourceArgs{
		RestApi:  gateway.ID(),
		PathPart: pulumi.String("{proxy+}"),
		ParentId: gateway.RootResourceId,
	})
	if err != nil {
		return nil, nil, err
	}

	// Add a method to the API Gateway.
	_, err = apigateway.NewMethod(ctx, "AnyMethod", &apigateway.MethodArgs{
		HttpMethod:    pulumi.String("ANY"),
		Authorization: pulumi.String("NONE"),
		RestApi:       gateway.ID(),
		ResourceId:    apiresource.ID(),
	})
	if err != nil {
		return nil, nil, err
	}

	// Add an integration to the API Gateway.
	// This makes communication between the API Gateway and the Lambda function work
	_, err = apigateway.NewIntegration(ctx, "LambdaIntegration", &apigateway.IntegrationArgs{
		HttpMethod:            pulumi.String("ANY"),
		IntegrationHttpMethod: pulumi.String("POST"),
		ResourceId:            apiresource.ID(),
		RestApi:               gateway.ID(),
		Type:                  pulumi.String("AWS_PROXY"),
		Uri:                   function.InvokeArn,
	})
	if err != nil {
		return nil, nil, err
	}

	// Add a resource based policy to the Lambda function.
	// This is the final step and allows AWS API Gateway to communicate with the AWS Lambda function
	permission, err := lambda.NewPermission(ctx, "APIPermission", &lambda.PermissionArgs{
		Action:    pulumi.String("lambda:InvokeFunction"),
		Function:  function.Name,
		Principal: pulumi.String("apigateway.amazonaws.com"),
		SourceArn: pulumi.Sprintf("arn:aws:execute-api:%s:%s:%s/*/*/*", region.Name, account.AccountId, gateway.ID()),
	}, pulumi.DependsOn([]pulumi.Resource{apiresource}))
	if err != nil {
		return nil, nil, err
	}

	// Create a new deployment
	_, err = apigateway.NewDeployment(ctx, "APIDeployment", &apigateway.DeploymentArgs{
		Description:      pulumi.String("UpperCase API deployment"),
		RestApi:          gateway.ID(),
		StageDescription: pulumi.String("Production"),
		StageName:        pulumi.String("prod"),
	}, pulumi.DependsOn([]pulumi.Resource{apiresource, function, permission}))
	if err != nil {
		return nil, nil, err
	}

	apiGwEndpointWithoutProtocol := pulumi.Sprintf("%s.execute-api.%s.amazonaws.com", gateway.ID(), region.Name)
	apiGwStageName := pulumi.String("prod")

	return &apiGwEndpointWithoutProtocol, &apiGwStageName, nil
}
