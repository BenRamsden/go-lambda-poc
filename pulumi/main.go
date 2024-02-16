package main

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		name := "jugo-go-lambda-poc"

		// TODO: Fetch root account zone id
		// TODO: Create poc.sandbox.jugo.io subdomain

		userPool, err := createCognito(ctx, name)
		if err != nil {
			return err
		}

		bucket, err := createBucket(ctx, name)
		if err != nil {
			return err
		}

		bucketOriginAccessIdentity, err := createBucketCloudfrontOrigin(ctx, name, bucket)
		if err != nil {
			return err
		}

		tables, err := createDynamo(ctx, name)
		if err != nil {
			return err
		}

		function, err := createLambda(ctx, name, tables.usersTable, tables.assetsTable, userPool)
		if err != nil {
			return err
		}

		apiGwEndpointWithoutProtocol, apiGwStageName, err := createApiGW(ctx, name, function)
		if err != nil {
			return err
		}

		dist, err := createCloudfront(ctx, name, bucket, bucketOriginAccessIdentity, apiGwEndpointWithoutProtocol, apiGwStageName)
		if err != nil {
			return err
		}

		ctx.Export("BucketName", bucket.ID())
		ctx.Export("APIGatewayURI", apiGwEndpointWithoutProtocol)
		ctx.Export("CloudfrontURI", dist.DomainName)
		ctx.Export("DynamoDBTables", pulumi.StringArray{tables.usersTable.Name, tables.assetsTable.Name})

		return nil
	})
}
