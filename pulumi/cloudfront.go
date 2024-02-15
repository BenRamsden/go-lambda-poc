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

func createCloudfront(ctx *pulumi.Context, name string, bucket *s3.BucketV2, bucketOriginAccessIdentity *cloudfront.OriginAccessIdentity, apiGwEndpointWithoutProtocol *pulumi.StringOutput, apiGwStageName *pulumi.String) (*cloudfront.Distribution, error) {
	OriginId := bucket.Arn

	dist, err := cloudfront.NewDistribution(ctx, name+"-cf-dist", &cloudfront.DistributionArgs{
		Origins: cloudfront.DistributionOriginArray{
			&cloudfront.DistributionOriginArgs{
				OriginId:   OriginId,
				DomainName: bucket.BucketRegionalDomainName,
				S3OriginConfig: cloudfront.DistributionOriginS3OriginConfigArgs{
					OriginAccessIdentity: bucketOriginAccessIdentity.CloudfrontAccessIdentityPath,
				},
			},
			&cloudfront.DistributionOriginArgs{
				OriginId:   apiGwEndpointWithoutProtocol,
				OriginPath: pulumi.Sprintf("/%s", apiGwStageName),
				DomainName: apiGwEndpointWithoutProtocol,
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
		//LoggingConfig: &cloudfront.DistributionLoggingConfigArgs{
		//	IncludeCookies: pulumi.Bool(false),
		//	Bucket:         pulumi.String("mylogs.s3.amazonaws.com"),
		//	Prefix:         pulumi.String("myprefix"),
		//},
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
			createAPIOrderedCacheBehaviour(apiGwEndpointWithoutProtocol, pulumi.String("/graphql")),
			createAPIOrderedCacheBehaviour(apiGwEndpointWithoutProtocol, pulumi.String("/playground")),
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
			CloudfrontDefaultCertificate: pulumi.Bool(true),
		},
	})

	if err != nil {
		return nil, err
	}
	return dist, nil
}
