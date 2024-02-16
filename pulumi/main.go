package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		name := "jugo-go-lambda-poc"

		// Stack reference to jugo-sandbox-base
		base, err := pulumi.NewStackReference(ctx, "jugo-sandbox-base", nil)
		if err != nil {
			return err
		}
		cloudFrontAcmCertArn := base.GetStringOutput(pulumi.String("cloudFrontAcmCertArn"))
		hostedZoneId := base.GetStringOutput(pulumi.String("zoneId"))

		poc := config.New(ctx, "jugo-go-lambda-poc")
		if err != nil {
			return err
		}
		targetUrl := pulumi.String(poc.Require("targetUrl"))

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
			cloudFrontAcmCertArn:         cloudFrontAcmCertArn,
			targetUrl:                    targetUrl,
		})
		if err != nil {
			return err
		}

		pocARecord, err := createARecord(ctx, name, &CreateARecordArgs{hostedZoneId: hostedZoneId, dist: dist, targetUrl: targetUrl})
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
