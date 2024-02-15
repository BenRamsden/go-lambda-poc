# jugo-go-lambda-poc load test

## Install

```
npm install -g artillery@latest
```

## Seed the database

```
artillery run assets-seed.yaml
```

## Run

```
artillery run assets-load-test.yaml
```

## TODO: Run distributed load test

https://github.com/artilleryio/artillery?tab=readme-ov-file#artillerycloud-scale-load-testing