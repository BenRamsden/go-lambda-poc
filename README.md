# jugo-go-lambda-poc

A proof of concept how Go Serverless Lambda can be utilized as a GraphQL backend for React.

# Setup

## Select Pulumi Stack

```
cd pulumi
sandbox
pulumi stack select jugo-go-lambda-poc
```

# Deploy

```
make package deploy
```