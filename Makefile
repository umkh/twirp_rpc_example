#!make
include .env
export $(shell sed 's/=.*//' .env)

generate:
	cd api; buf generate; cd ..

clean:
	rm -rf ./pkg/genproto; rm -rf ./pkg/swaggers

migrate-up:
	migrate -path ./internal/migrations -database 'postgres://$(PG_USER):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(PG_DATABASE)?sslmode=$(PG_SSL)' up

migrate-down:
	migrate -path ./internal/migrations -database 'postgres://$(PG_USER):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(PG_DATABASE)?sslmode=$(PG_SSL)' down