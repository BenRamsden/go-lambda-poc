package main

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		name := "jugo-go-lambda-poc"

		// Stack reference to jugo-sandbox-base
		base, err := pulumi.NewStackReference(ctx, "jugo-sandbox-base", nil)
		if err != nil {
			return err
		}

		acmCertArn := base.GetStringOutput(pulumi.String("cloudFrontAcmCertArn"))
		hostedZoneId := base.GetStringOutput(pulumi.String("zoneId"))

		bucket, err := createBucket(ctx, name)
		if err != nil {
			return err
		}

		bucketOriginAccessIdentity, err := createBucketCloudfrontOrigin(ctx, name, &CreateBucketCloudfrontOriginArgs{bucket: bucket})
		if err != nil {
			return err
		}

		tables, err := createDynamo(ctx, name)
		if err != nil {
			return err
		}

		function, err := createLambda(ctx, name, &CreateLambdaArgs{usersTable: tables.usersTable, assetsTable: tables.assetsTable})
		if err != nil {
			return err
		}

		apiGwEndpointWithoutProtocol, apiGwStageName, err := createApiGW(ctx, name, &CreateApiGWArgs{function: function})
		if err != nil {
			return err
		}

		dist, err := createCloudfront(ctx, name, &CreateCloudfrontArgs{
			bucket:                       bucket,
			bucketOriginAccessIdentity:   bucketOriginAccessIdentity,
			apiGwEndpointWithoutProtocol: apiGwEndpointWithoutProtocol,
			apiGwStageName:               apiGwStageName,
			acmCertArn:                   acmCertArn,
		})
		if err != nil {
			return err
		}

		pocARecord, err := createARecord(ctx, name, &CreateARecordArgs{hostedZoneId: hostedZoneId, dist: dist})
		if err != nil {
			return err
		}

		ctx.Export("BucketName", bucket.ID())
		ctx.Export("APIGatewayURI", apiGwEndpointWithoutProtocol)
		ctx.Export("CloudfrontURI", dist.DomainName)
		ctx.Export("DynamoDBTables", pulumi.StringArray{tables.usersTable.Name, tables.assetsTable.Name})
		ctx.Export("DomainName", pocARecord.Name)

		return nil
	})
}
