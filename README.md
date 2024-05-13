# go-lambda-poc

A proof of concept how Go Serverless Lambda can be utilized as a GraphQL backend for React.

# Pre-requisites

- Install Go
  - Mac `brew install go`
  - Manjaro `sudo pacman -Syu && sudo pacman -S go`
- Install Node & Yarn
  - Mac `brew install node && npm install -g yarn`
  - Manjaro `sudo pacman -Syu && sudo pacman -S nodejs && npm install -g yarn`

# Setup

```shell
make init
```

# Run Locally

> :warning: TODO: Make DynamoDB tables creation part of the `make start` process

```shell
make start
./scripts/create-dynamo-tables.sh
```

## Stop Locally

```shell
make stop
```

## Select Pulumi Stack

```
cd pulumi
sandbox
pulumi stack select go-lambda-poc
```

# Deploy

```
make package deploy
```
