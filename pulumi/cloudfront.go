package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createCloudfront(ctx *pulumi.Context, bucket *s3.BucketV2, bucketOriginAccessIdentity *cloudfront.OriginAccessIdentity) (*cloudfront.Distribution, error) {
	OriginId := bucket.Arn

	dist, err := cloudfront.NewDistribution(ctx, "s3Distribution", &cloudfront.DistributionArgs{
		Origins: cloudfront.DistributionOriginArray{
			&cloudfront.DistributionOriginArgs{
				OriginId:   OriginId,
				DomainName: bucket.BucketRegionalDomainName,
				S3OriginConfig: cloudfront.DistributionOriginS3OriginConfigArgs{
					OriginAccessIdentity: bucketOriginAccessIdentity.CloudfrontAccessIdentityPath,
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
		//Aliases: pulumi.StringArray{
		//	pulumi.String("mysite.example.com"),
		//	pulumi.String("yoursite.example.com"),
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
			ViewerProtocolPolicy: pulumi.String("allow-all"),
			MinTtl:               pulumi.Int(0),
			DefaultTtl:           pulumi.Int(3600),
			MaxTtl:               pulumi.Int(86400),
		},
		//OrderedCacheBehaviors: cloudfront.DistributionOrderedCacheBehaviorArray{
		//	&cloudfront.DistributionOrderedCacheBehaviorArgs{
		//		PathPattern: pulumi.String("/content/immutable/*"),
		//		AllowedMethods: pulumi.StringArray{
		//			pulumi.String("GET"),
		//			pulumi.String("HEAD"),
		//			pulumi.String("OPTIONS"),
		//		},
		//		CachedMethods: pulumi.StringArray{
		//			pulumi.String("GET"),
		//			pulumi.String("HEAD"),
		//			pulumi.String("OPTIONS"),
		//		},
		//		TargetOriginId: OriginId,
		//		ForwardedValues: &cloudfront.DistributionOrderedCacheBehaviorForwardedValuesArgs{
		//			QueryString: pulumi.Bool(false),
		//			Headers: pulumi.StringArray{
		//				pulumi.String("Origin"),
		//			},
		//			Cookies: &cloudfront.DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs{
		//				Forward: pulumi.String("none"),
		//			},
		//		},
		//		MinTtl:               pulumi.Int(0),
		//		DefaultTtl:           pulumi.Int(86400),
		//		MaxTtl:               pulumi.Int(31536000),
		//		Compress:             pulumi.Bool(true),
		//		ViewerProtocolPolicy: pulumi.String("redirect-to-https"),
		//	},
		//	&cloudfront.DistributionOrderedCacheBehaviorArgs{
		//		PathPattern: pulumi.String("/content/*"),
		//		AllowedMethods: pulumi.StringArray{
		//			pulumi.String("GET"),
		//			pulumi.String("HEAD"),
		//			pulumi.String("OPTIONS"),
		//		},
		//		CachedMethods: pulumi.StringArray{
		//			pulumi.String("GET"),
		//			pulumi.String("HEAD"),
		//		},
		//		TargetOriginId: OriginId,
		//		ForwardedValues: &cloudfront.DistributionOrderedCacheBehaviorForwardedValuesArgs{
		//			QueryString: pulumi.Bool(false),
		//			Cookies: &cloudfront.DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs{
		//				Forward: pulumi.String("none"),
		//			},
		//		},
		//		MinTtl:               pulumi.Int(0),
		//		DefaultTtl:           pulumi.Int(3600),
		//		MaxTtl:               pulumi.Int(86400),
		//		Compress:             pulumi.Bool(true),
		//		ViewerProtocolPolicy: pulumi.String("redirect-to-https"),
		//	},
		//},
		//PriceClass: pulumi.String("PriceClass_200"),
		Restrictions: &cloudfront.DistributionRestrictionsArgs{
			GeoRestriction: &cloudfront.DistributionRestrictionsGeoRestrictionArgs{
				RestrictionType: pulumi.String("whitelist"),
				Locations: pulumi.StringArray{
					//pulumi.String("US"),
					//pulumi.String("CA"),
					pulumi.String("GB"),
					//pulumi.String("DE"),
				},
			},
		},
		Tags: pulumi.StringMap{
			"Environment": pulumi.String("sandbox"),
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
