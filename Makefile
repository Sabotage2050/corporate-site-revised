init: ## Create environment variables
	cp .env.example .env
# App deployment build and run (debug mode with Delve)
development:
	BUILD_TARGET=development docker compose build
	docker compose up

# Test-debug build and run (for local test debugging)
test-debug:
	BUILD_TARGET=test-debug docker compose build
	docker compose up

# staging-test build and run (for local staging-test)
staging-debug:
	BUILD_TARGET=staging docker compose build
	docker compose up

# Test build and run without docker-compose (for CI)
ci:
	docker build --target ci -t ci-image .
	docker compose up

# Production deploy build (for GitHub Actions)
deploy:
	docker build --target deploy -t prod-image .

# Clean up Docker resources
clean:
	docker compose down --volumes --remove-orphans
	docker system prune -f

# sqlboilerでorm作成
sqlboiler:
	docker compose exec backend sqlboiler mysql

#oapicodegenで作成
PHONY: oapicodegen
oapicodegen:
	$(eval NAME := $(shell read -p "Enter oapicodegen name: " name; echo $$name))
	@mkdir -p backend/infra/api/$(NAME)
	@(oapi-codegen --config backend/oapi/config/$(NAME).yaml backend/oapi/$(NAME).yaml)

# Database migrations commands
db-create: ## Create migration files
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir database/migrations -seq create_$$name

db-up: ## Run all pending migrations
	migrate -path database/migrations \
		-database "mysql://root:password@tcp(localhost:3306)/manage_db" up

db-down: ## Rollback all migrations
	migrate -path database/migrations \
		-database "mysql://root:password@tcp(localhost:3306)/manage_db" down

db-reset: ## Drop all tables and rerun all migrations
	migrate -path database/migrations \
		-database "mysql://root:password@tcp(localhost:3306)/manage_db" drop -f
	migrate -path database/migrations \
		-database "mysql://root:password@tcp(localhost:3306)/manage_db" up

db-version: ## Show current migration version
	migrate -path database/migrations \
		-database "mysql://root:password@tcp(localhost:3306)/manage_db" version

db-force: ## Force set migration version
	@read -p "Enter version: " version; \
	migrate -path database/migrations \
		-database "mysql://root:password@tcp(localhost:3306)/manage_db" force $$version

db-up-step: ## Run N pending migrations forward
	@read -p "Enter number of migrations to run: " num; \
	migrate -path database/migrations \
		-database "mysql://root:password@tcp(localhost:3306)/manage_db" up $$num

db-down-step: ## Rollback N migrations
	@read -p "Enter number of migrations to rollback: " num; \
	migrate -path database/migrations \
		-database "mysql://root:password@tcp(localhost:3306)/manage_db" down $$num

