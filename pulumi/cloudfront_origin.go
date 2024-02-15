package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createBucketCloudfrontOrigin(ctx *pulumi.Context, name string, bucket *s3.BucketV2) (*cloudfront.OriginAccessIdentity, error) {
	bucketOriginAccessIdentity, err := cloudfront.NewOriginAccessIdentity(ctx, name+"-identity", &cloudfront.OriginAccessIdentityArgs{
		Comment: pulumi.String(name + "-identity"),
	})
	if err != nil {
		return nil, err
	}

	_, err = s3.NewBucketPolicy(ctx, name+"-policy", &s3.BucketPolicyArgs{
		Bucket: bucket.ID(),
		Policy: pulumi.Sprintf(`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "PublicReadGetObject",
            "Effect": "Allow",
            "Principal": {
				"AWS": "%s"
			},
            "Action": [
                "s3:GetObject"
            ],
            "Resource": [
				"arn:aws:s3:::%s/*"
            ]
        }
    ]
}`, bucketOriginAccessIdentity.IamArn, bucket.ID()),
	})
	if err != nil {
		return nil, err
	}

	return bucketOriginAccessIdentity, nil
}
