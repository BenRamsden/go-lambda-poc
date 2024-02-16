package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createAPIOrderedCacheBehaviour(apiGwEndpointWithoutProtocol *pulumi.StringOutput, pathPattern pulumi.String) *cloudfront.DistributionOrderedCacheBehaviorArgs {
	return &cloudfront.DistributionOrderedCacheBehaviorArgs{
		PathPattern: pathPattern,
		AllowedMethods: pulumi.StringArray{
			pulumi.String("DELETE"),
			pulumi.String("GET"),
			pulumi.String("HEAD"),
			pulumi.String("OPTIONS"),
			pulumi.String("PATCH"),
			pulumi.String("POST"),
			pulumi.String("PUT"),
		},
		CachedMethods: pulumi.StringArray{
			pulumi.String("GET"),
			pulumi.String("HEAD"),
		},
		TargetOriginId: apiGwEndpointWithoutProtocol,
		ForwardedValues: &cloudfront.DistributionOrderedCacheBehaviorForwardedValuesArgs{
			QueryString: pulumi.Bool(false),
			Headers: pulumi.StringArray{
				pulumi.String("Origin"),
				pulumi.String("Authorization"),
			},
			Cookies: &cloudfront.DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs{
				Forward: pulumi.String("none"),
			},
		},
		MinTtl:               pulumi.Int(0),
		DefaultTtl:           pulumi.Int(0),
		MaxTtl:               pulumi.Int(0),
		ViewerProtocolPolicy: pulumi.String("redirect-to-https"),
	}
}

type CreateCloudfrontArgs struct {
	frontendBucket               *s3.BucketV2
	loggingBucket                *s3.BucketV2
	bucketOriginAccessIdentity   *cloudfront.OriginAccessIdentity
	apiGwEndpointWithoutProtocol *pulumi.StringOutput
	apiGwStageName               *pulumi.String
	cloudFrontAcmCertArn         pulumi.StringOutput
	targetUrl                    pulumi.String
}

func createCloudfront(ctx *pulumi.Context, name string, args *CreateCloudfrontArgs) (*cloudfront.Distribution, error) {
	OriginId := args.frontendBucket.Arn

	dist, err := cloudfront.NewDistribution(ctx, name+"-cf-dist", &cloudfront.DistributionArgs{
		Origins: cloudfront.DistributionOriginArray{
			&cloudfront.DistributionOriginArgs{
				OriginId:   OriginId,
				DomainName: args.frontendBucket.BucketRegionalDomainName,
				S3OriginConfig: cloudfront.DistributionOriginS3OriginConfigArgs{
					OriginAccessIdentity: args.bucketOriginAccessIdentity.CloudfrontAccessIdentityPath,
				},
			},
			&cloudfront.DistributionOriginArgs{
				OriginId:   args.apiGwEndpointWithoutProtocol,
				OriginPath: pulumi.Sprintf("/%s", args.apiGwStageName),
				DomainName: args.apiGwEndpointWithoutProtocol,
				CustomOriginConfig: &cloudfront.DistributionOriginCustomOriginConfigArgs{
					HttpPort:             pulumi.Int(80),
					HttpsPort:            pulumi.Int(443),
					OriginProtocolPolicy: pulumi.String("https-only"),
					OriginSslProtocols: pulumi.StringArray{
						pulumi.String("TLSv1.2"),
					},
				},
			},
		},
		Enabled:           pulumi.Bool(true),
		IsIpv6Enabled:     pulumi.Bool(true),
		Comment:           pulumi.String("Some comment"),
		DefaultRootObject: pulumi.String("index.html"),
		LoggingConfig: &cloudfront.DistributionLoggingConfigArgs{
			IncludeCookies: pulumi.Bool(false),
			Bucket:         args.loggingBucket.BucketDomainName,
		},
		DefaultCacheBehavior: &cloudfront.DistributionDefaultCacheBehaviorArgs{
			AllowedMethods: pulumi.StringArray{
				pulumi.String("DELETE"),
				pulumi.String("GET"),
				pulumi.String("HEAD"),
				pulumi.String("OPTIONS"),
				pulumi.String("PATCH"),
				pulumi.String("POST"),
				pulumi.String("PUT"),
			},
			CachedMethods: pulumi.StringArray{
				pulumi.String("GET"),
				pulumi.String("HEAD"),
			},
			TargetOriginId: OriginId,
			ForwardedValues: &cloudfront.DistributionDefaultCacheBehaviorForwardedValuesArgs{
				QueryString: pulumi.Bool(false),
				Cookies: &cloudfront.DistributionDefaultCacheBehaviorForwardedValuesCookiesArgs{
					Forward: pulumi.String("none"),
				},
			},
			MinTtl:               pulumi.Int(0),
			DefaultTtl:           pulumi.Int(0),
			MaxTtl:               pulumi.Int(0),
			ViewerProtocolPolicy: pulumi.String("redirect-to-https"),
		},
		OrderedCacheBehaviors: cloudfront.DistributionOrderedCacheBehaviorArray{
			createAPIOrderedCacheBehaviour(args.apiGwEndpointWithoutProtocol, pulumi.String("/graphql")),
			createAPIOrderedCacheBehaviour(args.apiGwEndpointWithoutProtocol, pulumi.String("/playground")),
		},
		PriceClass: pulumi.String("PriceClass_200"),
		Restrictions: &cloudfront.DistributionRestrictionsArgs{
			GeoRestriction: &cloudfront.DistributionRestrictionsGeoRestrictionArgs{
				RestrictionType: pulumi.String("whitelist"),
				Locations: pulumi.StringArray{
					pulumi.String("GB"),
				},
			},
		},
		CustomErrorResponses: &cloudfront.DistributionCustomErrorResponseArray{
			&cloudfront.DistributionCustomErrorResponseArgs{
				ErrorCode:        pulumi.Int(404),
				ResponseCode:     pulumi.Int(200),
				ResponsePagePath: pulumi.String("/index.html"),
			},
		},
		ViewerCertificate: &cloudfront.DistributionViewerCertificateArgs{
			AcmCertificateArn:      args.cloudFrontAcmCertArn,
			SslSupportMethod:       pulumi.String("sni-only"),
			MinimumProtocolVersion: pulumi.String("TLSv1.2_2021"),
		},
		Aliases: pulumi.StringArray{
			args.targetUrl,
		},
	})

	if err != nil {
		return nil, err
	}
	return dist, nil
}

func createLoggingBucket(ctx *pulumi.Context, name string) (*s3.BucketV2, error) {
	loggingBucket, err := s3.NewBucketV2(ctx, name+"-logs", &s3.BucketV2Args{
		Bucket: pulumi.String(name + "-logs"),
	})
	if err != nil {
		return nil, err
	}
	_, err = s3.NewBucketOwnershipControls(ctx, name+"-logs-ownership", &s3.BucketOwnershipControlsArgs{
		Bucket: loggingBucket.ID(),
		Rule: &s3.BucketOwnershipControlsRuleArgs{
			ObjectOwnership: pulumi.String("BucketOwnerPreferred"),
		},
	})
	if err != nil {
		return nil, err
	}
	return loggingBucket, nil
}
