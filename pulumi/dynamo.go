package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createDynamoTable(ctx *pulumi.Context, name string) (*dynamodb.Table, error) {
	table, err := dynamodb.NewTable(ctx, name, &dynamodb.TableArgs{
		Name:     pulumi.String(name),
		HashKey:  pulumi.String("PK"),
		RangeKey: pulumi.String("SK"),
		Tags: pulumi.StringMap{
			"Environment": pulumi.String("sandbox"),
		},
		BillingMode: pulumi.String("PAY_PER_REQUEST"),
		Attributes: dynamodb.TableAttributeArray{
			&dynamodb.TableAttributeArgs{
				Name: pulumi.String("PK"),
				Type: pulumi.String("S"),
			},
			&dynamodb.TableAttributeArgs{
				Name: pulumi.String("SK"),
				Type: pulumi.String("S"),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return table, nil
}

// define return type struct including userTable assetsTable
type DynamoTables struct {
	usersTable  *dynamodb.Table
	assetsTable *dynamodb.Table
}

func createDynamo(ctx *pulumi.Context, name string) (*DynamoTables, error) {
	usersTable, err := createDynamoTable(ctx, name+"-users")
	if err != nil {
		return nil, err
	}
	assetsTable, err := createDynamoTable(ctx, name+"-assets")
	if err != nil {
		return nil, err
	}
	return &DynamoTables{
		usersTable:  usersTable,
		assetsTable: assetsTable,
	}, nil
}
