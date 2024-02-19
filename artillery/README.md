# jugo-go-lambda-poc load test

## Install

```
npm install -g artillery@latest
```

## Seed the database

```
export TOKEN="Bearer <token>"
artillery run assets-seed.yaml
```

## Run

```
export TOKEN="Bearer <token>"
artillery run assets-load-test.yaml
```

## Run in fargate
Count 15 is for 5000rps. 1 runner is good for ~500rps.

```
artillery run-fargate --region eu-west-1 assets-load-test.yml --count 15
```

## TODO: Run distributed load test

https://github.com/artilleryio/artillery?tab=readme-ov-file#artillerycloud-scale-load-testing