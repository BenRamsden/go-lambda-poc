package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"os"
	"strings"
)

func crawlDirectory(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	filePaths := []string{}
	for _, file := range files {
		filePath := dir + "/" + file.Name()
		stat, err := os.Stat(filePath)
		if err != nil {
			return nil, err
		}
		if stat.IsDir() {
			files_in_folder, err := crawlDirectory(filePath)
			if err != nil {
				return nil, err
			}
			filePaths = append(filePaths, files_in_folder...)
		}
		if stat.Mode().IsRegular() {
			filePaths = append(filePaths, filePath)
		}
	}
	return filePaths, nil
}

func createBucket(ctx *pulumi.Context) (*s3.BucketV2, *cloudfront.OriginAccessIdentity, error) {
	bucket, err := s3.NewBucketV2(ctx, "jugo-go-lambda-poc", &s3.BucketV2Args{
		Bucket: pulumi.String("jugo-go-lambda-poc"),
		Tags: pulumi.StringMap{
			"environment": pulumi.String("sandbox"),
		},
	})
	if err != nil {
		return nil, nil, err
	}

	//Allow public read access to all objects in the bucket
	_, err = s3.NewBucketPublicAccessBlock(ctx, "jugo-go-lambda-poc-public", &s3.BucketPublicAccessBlockArgs{
		Bucket:                bucket.ID(),
		BlockPublicAcls:       pulumi.Bool(false),
		IgnorePublicAcls:      pulumi.Bool(false),
		BlockPublicPolicy:     pulumi.Bool(false),
		RestrictPublicBuckets: pulumi.Bool(false),
	})
	if err != nil {
		return nil, nil, err
	}

	//Object Ownership, set ACLs enabled
	_, err = s3.NewBucketOwnershipControls(ctx, "jugo-go-lambda-poc-ownership", &s3.BucketOwnershipControlsArgs{
		Bucket: bucket.ID(),
		Rule: &s3.BucketOwnershipControlsRuleArgs{
			ObjectOwnership: pulumi.String("BucketOwnerPreferred"),
		},
	})

	// Set bucket acl to public-read
	_, err = s3.NewBucketAclV2(ctx, "jugo-go-lambda-poc-acl", &s3.BucketAclV2Args{
		Bucket: bucket.ID(),
		Acl:    pulumi.String("public-read"),
	})

	bucketOriginAccessIdentity, err := cloudfront.NewOriginAccessIdentity(ctx, "jugo-go-lambda-poc-identity", &cloudfront.OriginAccessIdentityArgs{
		Comment: pulumi.String("jugo-go-lambda-poc-identity"),
	})

	_, err = s3.NewBucketPolicy(ctx, "jugo-go-lambda-poc-policy", &s3.BucketPolicyArgs{
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

	files, err := crawlDirectory("../bin/ui")
	for _, file := range files {
		key := strings.Replace(file, "../bin/ui/", "", 1)
		mimeType := pulumi.String("")

		// TODO: replace with more advanced MIME, such as filetype lib
		if strings.HasSuffix(key, ".html") {
			mimeType = pulumi.String("text/html")
		} else if strings.HasSuffix(key, ".css") {
			mimeType = pulumi.String("text/css")
		} else if strings.HasSuffix(key, ".js") {
			mimeType = pulumi.String("application/javascript")
		}
		_, err := s3.NewBucketObjectv2(ctx, key, &s3.BucketObjectv2Args{
			Bucket:      bucket.ID(),
			Source:      pulumi.NewFileAsset(file),
			Acl:         pulumi.String("public-read"),
			ContentType: mimeType,
			Key:         pulumi.String(key),
		})
		if err != nil {
			return nil, nil, err
		}
	}

	// Export the name of the bucket
	ctx.Export("bucketName", bucket.ID())
	return bucket, bucketOriginAccessIdentity, nil
}
