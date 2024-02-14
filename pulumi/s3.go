package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"mime"
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

func createBucket(ctx *pulumi.Context) (*s3.Bucket, error) {
	bucket, err := s3.NewBucket(ctx, "jugo-go-lambda-poc", &s3.BucketArgs{
		Bucket: pulumi.String("jugo-go-lambda-poc"),
		Acl:    pulumi.String("private"),
		Tags: pulumi.StringMap{
			"environment": pulumi.String("sandbox"),
		},
		Website: &s3.BucketWebsiteArgs{
			IndexDocument: pulumi.String("index.html"),
			ErrorDocument: pulumi.String("index.html"),
		},
	})
	if err != nil {
		return nil, err
	}

	// Allow public read access to all objects in the bucket
	_, err = s3.NewBucketPublicAccessBlock(ctx, "jugo-go-lambda-poc-public", &s3.BucketPublicAccessBlockArgs{
		Bucket:                bucket.ID(),
		BlockPublicAcls:       pulumi.Bool(false),
		IgnorePublicAcls:      pulumi.Bool(false),
		BlockPublicPolicy:     pulumi.Bool(false),
		RestrictPublicBuckets: pulumi.Bool(false),
	})
	if err != nil {
		return nil, err
	}

	// Object Ownership, set ACLs enabled
	_, err = s3.NewBucketOwnershipControls(ctx, "jugo-go-lambda-poc-ownership", &s3.BucketOwnershipControlsArgs{
		Bucket: bucket.ID(),
		Rule: &s3.BucketOwnershipControlsRuleArgs{
			ObjectOwnership: pulumi.String("BucketOwnerPreferred"),
		},
	})

	files, err := crawlDirectory("../bin/ui")
	for _, file := range files {
		key := strings.Replace(file, "../bin/ui", "", 1)
		mimeType := mime.TypeByExtension(key)
		_, err := s3.NewBucketObject(ctx, key, &s3.BucketObjectArgs{
			Bucket:      bucket.ID(),
			Source:      pulumi.NewFileAsset(file),
			Acl:         pulumi.String("public-read"),
			ContentType: pulumi.String(mimeType),
			Key:         pulumi.String(key),
		})
		if err != nil {
			return nil, err
		}
	}

	// Export the name of the bucket
	ctx.Export("bucketName", bucket.ID())
	return bucket, nil
}
