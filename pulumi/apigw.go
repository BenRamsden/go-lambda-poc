package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createApiGWRoute(ctx *pulumi.Context, name string, pathPart pulumi.String, gateway *apigateway.RestApi, function *lambda.Function) (*apigateway.Integration, error) {
	account, err := aws.GetCallerIdentity(ctx, nil)
	if err != nil {
		return nil, err
	}

	region, err := aws.GetRegion(ctx, &aws.GetRegionArgs{})
	if err != nil {
		return nil, err
	}

	resource, err := apigateway.NewResource(ctx, name+"-api-gw-resource", &apigateway.ResourceArgs{
		RestApi:  gateway.ID(),
		PathPart: pathPart,
		ParentId: gateway.RootResourceId,
	})
	if err != nil {
		return nil, err
	}

	method, err := apigateway.NewMethod(ctx, name+"-api-gw-any-method", &apigateway.MethodArgs{
		HttpMethod:    pulumi.String("ANY"),
		Authorization: pulumi.String("NONE"),
		RestApi:       gateway.ID(),
		ResourceId:    resource.ID(),
	}, pulumi.DependsOn([]pulumi.Resource{resource}))
	if err != nil {
		return nil, err
	}

	integration, err := apigateway.NewIntegration(ctx, name+"-api-gw-any-method-lambda-integration", &apigateway.IntegrationArgs{
		HttpMethod:            pulumi.String("ANY"),
		IntegrationHttpMethod: pulumi.String("ANY"),
		ResourceId:            resource.ID(),
		RestApi:               gateway.ID(),
		Type:                  pulumi.String("AWS_PROXY"),
		Uri:                   function.InvokeArn,
	}, pulumi.DependsOn([]pulumi.Resource{method}))
	if err != nil {
		return nil, err
	}

	_, err = lambda.NewPermission(ctx, name+"-api-gw-api-permission", &lambda.PermissionArgs{
		Action:    pulumi.String("lambda:InvokeFunction"),
		Function:  function.Name,
		Principal: pulumi.String("apigateway.amazonaws.com"),
		SourceArn: pulumi.Sprintf("arn:aws:execute-api:%s:%s:%s/*/*/*", region.Name, account.AccountId, gateway.ID()),
	}, pulumi.DependsOn([]pulumi.Resource{integration}))
	if err != nil {
		return nil, err
	}

	return integration, nil
}

type CreateApiGWArgs struct {
	goFunction *lambda.Function
	tsFunction *lambda.Function
}

func createApiGW(ctx *pulumi.Context, name string, args *CreateApiGWArgs) (*pulumi.StringOutput, *pulumi.String, error) {
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

	go_integration, err := createApiGWRoute(ctx, name+"-graphql", pulumi.String("graphql"), gateway, args.goFunction)
	if err != nil {
		return nil, nil, err
	}

	ts_integration, err := createApiGWRoute(ctx, name+"-graphql2", pulumi.String("graphql2"), gateway, args.tsFunction)
	if err != nil {
		return nil, nil, err
	}

	_, err = apigateway.NewDeployment(ctx, name+"-api-gw-deployment", &apigateway.DeploymentArgs{
		RestApi:   gateway.ID(),
		StageName: pulumi.String("prod"),
	}, pulumi.DependsOn([]pulumi.Resource{go_integration, ts_integration}))

	apiGwEndpointWithoutProtocol := pulumi.Sprintf("%s.execute-api.%s.amazonaws.com", gateway.ID(), region.Name)
	apiGwStageName := pulumi.String("prod")

	return &apiGwEndpointWithoutProtocol, &apiGwStageName, nil
}
