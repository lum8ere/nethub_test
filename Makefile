MIGRATIONS_DIR := ./migrations
DB_URL ?= postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable
BACKEND_DIR := ./backend
	
.PHONY: migrate-create migrate-up migrate-down migrate-rollback migrate-force db-up db-down swagger

migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Ошибка: Укажите имя миграции. Пример: make migrate-create name=init_schema"; \
		exit 1; \
	fi
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) up

migrate-down:
	@read -p "ВНИМАНИЕ! Это удалит ВСЕ таблицы. Вы уверены? [y/N]: " ans; \
	if [ "$$ans" = "y" ] || [ "$$ans" = "Y" ]; then \
		migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) down; \
	else \
		echo "Отменено."; \
	fi

migrate-rollback:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) down 1

migrate-force:
	@if [ -z "$(version)" ]; then \
		echo "Ошибка: Укажите версию. Пример: make migrate-force version=1"; \
		exit 1; \
	fi
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) force $(version)

swagger:
	cd backend && swag init -g ./cmd/api/main.go --parseDependency --parseInternal

.PHONY: test test-v test-cover test-race cover-html

test:
	@echo "Running tests..."
	cd $(BACKEND_DIR) && go test ./...

test-v:
	@echo "Running tests in verbose mode..."
	cd $(BACKEND_DIR) && go test -v ./...

test-cover:
	@echo "Running tests with coverage..."
	cd $(BACKEND_DIR) && go test -cover ./...

cover-html: 
	cd $(BACKEND_DIR) && go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out