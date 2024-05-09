package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		name := "go-lambda-poc"

		// Stack reference to jugo-sandbox-base
		base, err := pulumi.NewStackReference(ctx, "jugo-sandbox-base", nil)
		if err != nil {
			return err
		}
		cloudFrontAcmCertArn := base.GetStringOutput(pulumi.String("cloudFrontAcmCertArn"))
		hostedZoneId := base.GetStringOutput(pulumi.String("zoneId"))

		poc := config.New(ctx, "go-lambda-poc")
		if err != nil {
			return err
		}
		targetUrl := pulumi.String(poc.Require("targetUrl"))

		frontendBucket, err := createFrontendBucket(ctx, name)
		if err != nil {
			return err
		}

		loggingBucket, err := createLoggingBucket(ctx, name)
		if err != nil {
			return err
		}

		bucketOriginAccessIdentity, err := createBucketCloudfrontOrigin(ctx, name, &CreateBucketCloudfrontOriginArgs{bucket: frontendBucket})
		if err != nil {
			return err
		}

		tables, err := createDynamo(ctx, name)
		if err != nil {
			return err
		}

		function, err := createLambda(ctx, name, &CreateLambdaArgs{
			usersTable:    tables.usersTable,
			assetsTable:   tables.assetsTable,
			Runtime:       pulumi.String("provided.al2023"),
			Code:          pulumi.NewFileArchive("../bin/lambda/api/api.zip"),
			Architectures: pulumi.StringArray{pulumi.String("arm64")},
			sentryDsn:     pulumi.String(poc.Require("sentryDsn")),
		})
		if err != nil {
			return err
		}

		// TODO: Set tsFunction to typescript lambda function
		apiGwEndpointWithoutProtocol, apiGwStageName, err := createApiGW(ctx, name, &CreateApiGWArgs{goFunction: function, tsFunction: function})
		if err != nil {
			return err
		}

		dist, err := createCloudfront(ctx, name, &CreateCloudfrontArgs{
			frontendBucket:               frontendBucket,
			loggingBucket:                loggingBucket,
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

		ctx.Export("BucketName", frontendBucket.ID())
		ctx.Export("APIGatewayURI", apiGwEndpointWithoutProtocol)
		ctx.Export("CloudfrontURI", dist.DomainName)
		ctx.Export("DynamoDBTables", pulumi.StringArray{tables.usersTable.Name, tables.assetsTable.Name})
		ctx.Export("DomainName", pocARecord.Name)

		return nil
	})
}
