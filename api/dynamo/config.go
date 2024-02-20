package dynamo

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-xray-sdk-go/instrumentation/awsv2"
)

var (
	assetsTableName string = "jugo-go-lambda-poc-assets"
	// usersTableName  string = "jugo-go-lambda-poc-users"
)

func init() {
	envAssetsTableName := os.Getenv("ASSETS_TABLE_NAME")
	if envAssetsTableName != "" {
		assetsTableName = envAssetsTableName
	}

	// envUsersTableName := os.Getenv("USERS_TABLE_NAME")
	// if envUsersTableName != "" {
	// 	usersTableName = envUsersTableName
	// }
}

// Targets dynamodb-local at http://localhost:8000
func NewFromLocal() Repository {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000"}, nil
			})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "dummy", SecretAccessKey: "dummy", SessionToken: "dummy",
				Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	svc := dynamodb.NewFromConfig(cfg)

	return New(svc)
}

// NewFromEnv creates a new DynamoDB client from the environment.
func NewFromEnv() Repository {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	awsv2.AWSV2Instrumentor(&cfg.APIOptions)
	
	// Using the Config value, create the DynamoDB client
	svc := dynamodb.NewFromConfig(cfg)

	return New(svc)
}
