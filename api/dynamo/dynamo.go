package dynamo

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/jugo-io/go-poc/api/model"
)

type Repository interface {
	model.AssetRepository
}

type repository struct {
	*dynamodb.Client
}

func New(client *dynamodb.Client) Repository {
	if client == nil {
		panic("nil dynamodb client")
	}
	resp, err := client.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(5),
	})
	if err != nil {
		panic(err)
	}
	for _, table := range resp.TableNames {
		fmt.Println(table)
	}

	return &repository{client}
}

// CreateAsset implements Repository.
func (repo *repository) CreateAsset(asset model.Asset) (model.Asset, error) {
	dbAsset := AssetFromModel(asset)

	item, err := attributevalue.MarshalMap(dbAsset)
	if err != nil {
		return model.Asset{}, err
	}

	fmt.Println(item)

	_, err = repo.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(assetsTableName),
		Item:      item,
	})

	return asset, err
}

// GetAssets implements Repository.
func (repo *repository) GetAssets(ownerId string) ([]model.Asset, error) {
	var err error
	var dbAssets []Asset
	var response *dynamodb.QueryOutput

	keyEx := expression.Key("PK").Equal(expression.Value(ownerId))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		return nil, err
	}

	queryPaginator := dynamodb.NewQueryPaginator(repo, &dynamodb.QueryInput{
		TableName:                 aws.String(assetsTableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})

	for queryPaginator.HasMorePages() {
		response, err = queryPaginator.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}

		var dbAssetsPage []Asset
		err = attributevalue.UnmarshalListOfMaps(response.Items, &dbAssetsPage)
		if err != nil {
			fmt.Println("Error unmarshalling:", err)
			return nil, err
		}

		dbAssets = append(dbAssets, dbAssetsPage...)
	}

	assets := make([]model.Asset, len(dbAssets))
	for i, dbAsset := range dbAssets {
		assets[i] = ModelFromAsset(dbAsset)
	}

	return assets, nil
}