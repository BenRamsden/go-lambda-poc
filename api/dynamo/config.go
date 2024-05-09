package dynamo

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-xray-sdk-go/instrumentation/awsv2"
)

func defaultTables() tables {
	var assetsTableName string = "go-lambda-poc-assets"
	envAssetsTableName := os.Getenv("ASSETS_TABLE_NAME")
	if envAssetsTableName != "" {
		assetsTableName = envAssetsTableName
	}

	return tables{
		assets: assetsTableName,
	}
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

	return New(svc, defaultTables())
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

	return New(svc, defaultTables())
}

func NewForTest(id string) Repository {
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

	tables := tables{
		assets: fmt.Sprintf("assets-%s", id),
	}

	_, err = svc.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		TableName:   aws.String(tables.assets),
		BillingMode: types.BillingModePayPerRequest,
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("PK"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("SK"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("PK"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("SK"),
				KeyType:       types.KeyTypeRange,
			},
		},
		Tags: []types.Tag{
			{
				Key:   aws.String("Environment"),
				Value: aws.String("sandbox"),
			},
		},
	})

	if err != nil {
		log.Fatalf("unable to create table: %v", err)
	}

	return New(svc, tables)
}

func CleanupForTest(repo Repository) {
	dyn := repo.(*repository)
	if dyn == nil {
		return
	}

	fmt.Printf("Cleaning up test table: %s\n", dyn.tables.assets)
	_, err := dyn.DeleteTable(context.TODO(), &dynamodb.DeleteTableInput{
		TableName: aws.String(dyn.tables.assets),
	})
	if err != nil {
		log.Fatalf("unable to delete table: %v", err)
	}
}
