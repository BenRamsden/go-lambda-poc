package dynamo

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/jugo-io/go-poc/api/model"
)

type Repository interface {
	model.AssetRepository
}

type tables struct {
	assets string
}

type repository struct {
	*dynamodb.Client
	tables tables
}

func New(client *dynamodb.Client, tables tables) Repository {
	if client == nil {
		panic("nil dynamodb client")
	}

	return &repository{
		Client: client,
		tables: tables,
	}
}
