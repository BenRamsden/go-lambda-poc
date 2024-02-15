package main

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, bucketOriginAccessIdentity, err := createBucket(ctx)
		if err != nil {
			return err
		}

		tables, err := createDynamo(ctx)
		if err != nil {
			return err
		}

		apiGwEndpointWithoutProtocol, apiGwStageName, err := createLambdas(ctx, tables.usersTable, tables.assetsTable)
		if err != nil {
			return err
		}

		dist, err := createCloudfront(ctx, bucket, bucketOriginAccessIdentity, apiGwEndpointWithoutProtocol, apiGwStageName)
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
