#!/bin/bash
set -euo pipefail

CreateGenericDynamoDBTable() {
  AWS_ACCESS_KEY_ID=X AWS_SECRET_ACCESS_KEY=X aws dynamodb create-table \
      --table-name "$1" \
      --attribute-definitions \
          AttributeName=PK,AttributeType=S \
          AttributeName=SK,AttributeType=S \
      --key-schema \
          AttributeName=PK,KeyType=HASH \
          AttributeName=SK,KeyType=RANGE \
      --provisioned-throughput \
          ReadCapacityUnits=1,WriteCapacityUnits=1 \
      --endpoint-url http://localhost:8000
          
}

CreateGenericDynamoDBTable "jugo-go-lambda-poc-assets"
CreateGenericDynamoDBTable "jugo-go-lambda-poc-users"