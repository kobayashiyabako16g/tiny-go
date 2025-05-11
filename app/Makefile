# Default Go BUILD
GOOS   ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
CGO_ENABLED ?= 0
# Database
DB_FILE = ./db/app.db
SCHEMA_FILE = ./db/schema.sql

# ==========================
# Project overall
# ==========================

.DEFAULT_GOAL := help
.PHONY: help
help: ## Show help messages
	@echo ''
	@grep -E '^[%/0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-22s\033[0m %s\n", $$1, $$2}'
	@echo ''

# ==========================
# Build
# ==========================

.PHONY: build
build: ## Build app (ex: make build GOOS=linux GOARCH=amd64)
	@echo "Building for GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED)"
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) go build -tags timetzdata -o ./bin/app ./cmd/*

.PHONY: clean
clean:
	rm -f bin/*

# ==========================
# Dev
# ==========================

.PHONY: test
test: ## Test (go test)
	go test $(shell go list ${MAKEFILE_DIR}/...)

.PHONY: vet
vet: ## Vet (go vet)
	go vet ./...

# ==========================
# Database
# ==========================
.PHONY: initSQLite
initSQLite: ## Init SQLite3 (dependency sqlite cli)
	sqlite3 $(DB_FILE) < $(SCHEMA_FILE)
	@echo "✅ Applied schema.sql to SQLite DB '$(DB_FILE)'."

.PHONY: initPostgres
initPostgres: ## Init PostgreSQL (requires psql CLI)
	spsql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d $(DB_NAME) -f $(SCHEMA_FILE)
	@echo "✅ Applied schema.sql to PostgreSQL DB '$(DB_NAME)'."
