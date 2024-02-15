package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createLambda(ctx *pulumi.Context, name string, usersTable *dynamodb.Table, assetsTable *dynamodb.Table) (*lambda.Function, error) {
	// Create an IAM role.
	role, err := iam.NewRole(ctx, name+"-task-exec-role", &iam.RoleArgs{
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
		return nil, err
	}

	// Attach a policy to allow writing logs to CloudWatch
	logPolicy, err := iam.NewRolePolicy(ctx, name+"-lambda-log-policy", &iam.RolePolicyArgs{
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

	dynamoPolicy, err := iam.NewRolePolicy(ctx, name+"-lambda-dynamo-policy", &iam.RolePolicyArgs{
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

	// Create the lambda using the args.
	function, err := lambda.NewFunction(
		ctx,
		name+"-lambda",
		&lambda.FunctionArgs{
			Name:    pulumi.String(name + "-lambda"),
			Handler: pulumi.String("handler"),
			Role:    role.Arn,
			Runtime: pulumi.String("provided.al2023"),
			Code:    pulumi.NewFileArchive("../bin/lambda/api/api.zip"),
			// Arm64
			Architectures: pulumi.StringArray{
				pulumi.String("arm64"),
			},
			Environment: lambda.FunctionEnvironmentArgs{
				Variables: pulumi.StringMap{
					"USERS_TABLE_NAME":  usersTable.Name,
					"ASSETS_TABLE_NAME": assetsTable.Name,
				},
			},
		},
		pulumi.DependsOn([]pulumi.Resource{logPolicy, dynamoPolicy}),
	)
	if err != nil {
		return nil, err
	}

	return function, nil
}
