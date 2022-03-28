# Mytheresa Promotions Test
## Challenge

[![N|Solid](https://cldup.com/dTxpPi9lDf.thumb.png)](https://nodesource.com/products/nsolid)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

This is my solution for Capitol offer challenge.

- Builded With
    - Golang REST API
    - Symfony REST API
    - Hexagonal Architecture
    - TDD, BDD
    - Docker, Docker compose
    - Unit tests

## Bussines Logic:

- You must take into account that this list could grow to have 20.000 products
- The prices are integers for example, 100.00€ would be 10000 .
- You can store the products as you see fit (json file, in memory, rdbms of choice)
- Products in the boots category have a 30% discount
- The product with sku = 000003 has a 15% discount
- When multiple discounts collide, the bigger discount must be applied.
- Can be filtered by category as a query string parameter
- (optional) Can be filtered by priceLessThan as a query string parameter, this filter applies before discounts are applied and will show products with prices
lesser than or equal the value provided.
- Returns a list of Product with the given discounts applied when necessary
- Must return at most 5 elements. (The order does not matter)


## How to deploy services

The services are conformed by a REST API in Golang and other implementation with Symfony:

- Just run the following command if you have docker and docker-compose installed:
```sh
    docker-compose up
```
- The port for golang implementation is 4000 and for Symfony 8000
- The endpoint for Golang implementation is:
```sh
    localhost:4000/v1/products
```
- The endpoint for Symfony implementation is:
```sh
    localhost:8000/products/list
```
- For run tests you must to run locally for Golang implementation you can run:
```sh
    go test -v ./...
```
- For run tests with Symfony implementation:
```sh
    composer run-script test
```


## License

MIT