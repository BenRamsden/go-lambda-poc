package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CreateApiGWArgs struct {
	goFunction *lambda.Function
	tsFunction *lambda.Function
}

func createApiGW(ctx *pulumi.Context, name string, args *CreateApiGWArgs) (*pulumi.StringOutput, *pulumi.String, error) {
	account, err := aws.GetCallerIdentity(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	region, err := aws.GetRegion(ctx, &aws.GetRegionArgs{})
	if err != nil {
		return nil, nil, err
	}

	// Create a new API Gateway.
	gateway, err := apigateway.NewRestApi(ctx, name+"-api-gw", &apigateway.RestApiArgs{
		Name: pulumi.String(name + "-api-gw"),
		// TODO: Narrow this down to just cloudfront
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

	// Go Lambda API Resource
	go_api_resource, err := apigateway.NewResource(ctx, name+"-api-gw-graphql-resource", &apigateway.ResourceArgs{
		RestApi:  gateway.ID(),
		PathPart: pulumi.String("graphql"),
		ParentId: gateway.RootResourceId,
	})
	if err != nil {
		return nil, nil, err
	}
	_, err = apigateway.NewMethod(ctx, name+"-api-gw-graphql-any-method", &apigateway.MethodArgs{
		HttpMethod:    pulumi.String("ANY"),
		Authorization: pulumi.String("NONE"),
		RestApi:       gateway.ID(),
		ResourceId:    go_api_resource.ID(),
	})
	if err != nil {
		return nil, nil, err
	}
	_, err = apigateway.NewIntegration(ctx, name+"api-gw-graphql-any-method-lambda-integration", &apigateway.IntegrationArgs{
		HttpMethod:            pulumi.String("ANY"),
		IntegrationHttpMethod: pulumi.String("ANY"),
		ResourceId:            go_api_resource.ID(),
		RestApi:               gateway.ID(),
		Type:                  pulumi.String("AWS_PROXY"),
		Uri:                   args.goFunction.InvokeArn,
	}, pulumi.DependsOn([]pulumi.Resource{go_api_resource}))
	if err != nil {
		return nil, nil, err
	}
	_, err = lambda.NewPermission(ctx, name+"api-gw-graphql-api-permission", &lambda.PermissionArgs{
		Action:    pulumi.String("lambda:InvokeFunction"),
		Function:  args.goFunction.Name,
		Principal: pulumi.String("apigateway.amazonaws.com"),
		SourceArn: pulumi.Sprintf("arn:aws:execute-api:%s:%s:%s/*/*/*", region.Name, account.AccountId, gateway.ID()),
	}, pulumi.DependsOn([]pulumi.Resource{go_api_resource}))
	if err != nil {
		return nil, nil, err
	}

	// TypeScript Lambda API Resource
	ts_api_resource, err := apigateway.NewResource(ctx, name+"-api-gw-graphql2-resource", &apigateway.ResourceArgs{
		RestApi:  gateway.ID(),
		PathPart: pulumi.String("graphql2"),
		ParentId: gateway.RootResourceId,
	})
	if err != nil {
		return nil, nil, err
	}
	_, err = apigateway.NewMethod(ctx, name+"-api-gw-graphql2-any-method", &apigateway.MethodArgs{
		HttpMethod:    pulumi.String("ANY"),
		Authorization: pulumi.String("NONE"),
		RestApi:       gateway.ID(),
		ResourceId:    ts_api_resource.ID(),
	})
	if err != nil {
		return nil, nil, err
	}
	_, err = apigateway.NewIntegration(ctx, name+"api-gw-graphql2-any-method-lambda-integration", &apigateway.IntegrationArgs{
		HttpMethod:            pulumi.String("ANY"),
		IntegrationHttpMethod: pulumi.String("ANY"),
		ResourceId:            ts_api_resource.ID(),
		RestApi:               gateway.ID(),
		Type:                  pulumi.String("AWS_PROXY"),
		Uri:                   args.tsFunction.InvokeArn,
	}, pulumi.DependsOn([]pulumi.Resource{ts_api_resource}))
	if err != nil {
		return nil, nil, err
	}
	_, err = lambda.NewPermission(ctx, name+"api-gw-graphql2-api-permission", &lambda.PermissionArgs{
		Action:    pulumi.String("lambda:InvokeFunction"),
		Function:  args.tsFunction.Name,
		Principal: pulumi.String("apigateway.amazonaws.com"),
		SourceArn: pulumi.Sprintf("arn:aws:execute-api:%s:%s:%s/*/*/*", region.Name, account.AccountId, gateway.ID()),
	}, pulumi.DependsOn([]pulumi.Resource{ts_api_resource}))
	if err != nil {
		return nil, nil, err
	}

	// Create a new deployment
	_, err = apigateway.NewDeployment(ctx, name+"api-deployment", &apigateway.DeploymentArgs{
		RestApi:          gateway.ID(),
		StageDescription: pulumi.String("Production"),
		StageName:        pulumi.String("prod"),
	}, pulumi.DependsOn([]pulumi.Resource{go_api_resource, ts_api_resource}))
	if err != nil {
		return nil, nil, err
	}

	apiGwEndpointWithoutProtocol := pulumi.Sprintf("%s.execute-api.%s.amazonaws.com", gateway.ID(), region.Name)
	apiGwStageName := pulumi.String("prod")

	return &apiGwEndpointWithoutProtocol, &apiGwStageName, nil
}
