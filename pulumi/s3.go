package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createBucket(ctx *pulumi.Context) (*s3.Bucket, error) {
	bucket, err := s3.NewBucket(ctx, "jugo-go-lambda-poc", &s3.BucketArgs{
		Bucket: pulumi.String("jugo-go-lambda-poc"),
		Acl:    pulumi.String("private"),
		Tags: pulumi.StringMap{
			"environment": pulumi.String("sandbox"),
		},
	})
	if err != nil {
		return nil, err
	}

	// Export the name of the bucket
	ctx.Export("bucketName", bucket.ID())
	return bucket, nil
}
