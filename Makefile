.PHONY: build-development build-staging build-production
.PHONY: start-development start-staging start-production
.PHONY: stop-development stop-staging stop-production

build-development:
	docker compose -f ./docker/development/docker-compose.yml  build --no-cache

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

