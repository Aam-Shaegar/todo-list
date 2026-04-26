include .env
export

export PROJECT_ROOT=$(shell pwd)
env-up:
	docker compose up -d todo-app-postgres

env-down:
	docker compose down todo-app-postgres

env-cleanup:
	@read -p "Clean all volume env files? [y/n]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todo-app-postgres && \
		rm -rf out/pgdata && \
		echo "Environment cleaned up."; \
	else \
		echo "Cleanup aborted."; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Error: Migration name is required. Usage: make migrate-create seq=<migration_name>"; \
		exit 1; \
	fi; \
	docker compose run --rm todo-app-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	make migrate-action action=up

migrate-down:
	make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Error: Migration action is required. Usage: make migrate-action action=<action>"; \
		exit 1; \
	fi; \

	docker compose run --rm todo-app-postgres-migrate \
		-path /migrations \
		-database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@todo-app-postgres:5432/$(POSTGRES_DB)?sslmode=disable" \
		"$(action)"

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder