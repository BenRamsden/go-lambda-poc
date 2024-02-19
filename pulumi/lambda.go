package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CreateLambdaArgs struct {
	usersTable    *dynamodb.Table
	assetsTable   *dynamodb.Table
	Runtime       pulumi.StringInput
	Code          pulumi.ArchiveInput
	Architectures pulumi.StringArrayInput
	sentryDsn     pulumi.StringInput
}

func createLambda(ctx *pulumi.Context, name string, args *CreateLambdaArgs) (*lambda.Function, error) {
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
						"dynamodb:PutItem",
						"dynamodb:Query"
					],
					"Resource": [
						"%s",	
						"%s"	
					]
				}
			]
		}`, args.assetsTable.Arn, args.usersTable.Arn),
	})

	// Create the lambda using the args.
	function, err := lambda.NewFunction(
		ctx,
		name+"-lambda",
		&lambda.FunctionArgs{
			Name:          pulumi.String(name + "-lambda"),
			Handler:       pulumi.String("handler"),
			Role:          role.Arn,
			Runtime:       args.Runtime,
			Code:          args.Code,
			Architectures: args.Architectures,
			Environment: lambda.FunctionEnvironmentArgs{
				Variables: pulumi.StringMap{
					"USERS_TABLE_NAME":  args.usersTable.Name,
					"ASSETS_TABLE_NAME": args.assetsTable.Name,
					// TODO: change to https://graphql.sandbox.jugo.io/graphql
					"AUTH0_AUDIENCE":     pulumi.String("http://localhost:4000/graphql"),
					"AUTH0_DOMAIN":       pulumi.String("https://auth.sandbox.jugo.io/"),
					"GIN_MODE":           pulumi.String("release"),
					"SENTRY_DSN":         args.sentryDsn,
					"SENTRY_ENVIRONMENT": pulumi.String("sandbox"),
					"SENTRY_RELEASE":     pulumi.String("0.0.0"),
				},
			},
			Timeout: pulumi.Int(30),
		},
		pulumi.DependsOn([]pulumi.Resource{logPolicy, dynamoPolicy}),
	)
	if err != nil {
		return nil, err
	}

	return function, nil
}
