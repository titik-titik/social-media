# CI/CD docker command
build-development:
	docker compose -f ./docker/development/docker-compose.yml build --no-cache

start-development:
	docker compose -f ./docker/development/docker-compose.yml up -d

stop-development:
	docker compose -f ./docker/development/docker-compose.yml down

build-staging:
	docker compose  -f ./docker/staging/docker-compose.yml build --no-cache

start-staging:
	docker compose -f ./docker/staging/docker-compose.yml up -d

stop-staging:
	docker compose -f ./docker/staging/docker-compose.yml down

build-production:
	docker compose  -f ./docker/production/docker-compose.yml build --no-cache

start-production:
	docker compose -f ./docker/production/docker-compose.yml up -d

stop-production:
	docker compose -f ./docker/production/docker-compose.yml down

# start local development
start-local:
	go run cmd/web/main.go

start-local-env:
	 GATEWAY_HOST=127.0.0.1 GATEWAY_APP_PORT=8001 go run cmd/web/main.go

# start all test
start-test:
	go test -v -count=1 ./test

start-test-env:
	 GATEWAY_HOST=127.0.0.1 GATEWAY_APP_PORT=8001 go test -v ./test/*
