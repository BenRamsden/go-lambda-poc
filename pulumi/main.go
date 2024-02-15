package main

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, bucketOriginAccessIdentity, err := createBucket(ctx)
		if err != nil {
			return err
		}

		apiGwEndpointWithoutProtocol, apiGwStageName, err := createLambdas(ctx)
		if err != nil {
			return err
		}

		dist, err := createCloudfront(ctx, bucket, bucketOriginAccessIdentity, apiGwEndpointWithoutProtocol, apiGwStageName)
		if err != nil {
			return err
		}

		pulumi.Printf("Bucket name: %s\n", bucket.ID())
		pulumi.Printf("Invocation URL: https://%s\n", apiGwEndpointWithoutProtocol)
		pulumi.Printf("Website URL: http://%s\n", dist.DomainName)

		return nil
	})
}
