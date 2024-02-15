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