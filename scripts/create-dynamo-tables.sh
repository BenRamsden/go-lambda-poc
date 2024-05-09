#!/bin/bash

set -euo pipefail

export AWS_PAGER="" # Disable the AWS CLI pager

CreateGenericDynamoDBTable() {
  response=$(AWS_ACCESS_KEY_ID=X AWS_SECRET_ACCESS_KEY=X aws dynamodb list-tables \
      --output text \
      --endpoint-url http://localhost:8000)

  if [[ $response == *"$1"* ]]; then
    AWS_ACCESS_KEY_ID=X AWS_SECRET_ACCESS_KEY=X aws dynamodb delete-table \
        --table-name "$1" \
        --endpoint-url http://localhost:8000
  fi

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

CreateGenericDynamoDBTable "go-lambda-poc-assets"
CreateGenericDynamoDBTable "go-lambda-poc-users"