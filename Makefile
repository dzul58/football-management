.PHONY: help run build test clean migrate-up migrate-down seed

help: ## Menampilkan bantuan
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

run: ## Menjalankan aplikasi
	go run cmd/api/main.go

build: ## Build aplikasi
	go build -o bin/api cmd/api/main.go

test: ## Menjalankan test
	go test -v ./...

clean: ## Membersihkan binary dan cache
	go clean
	rm -rf bin/

install: ## Install dependencies
	go mod download
	go mod tidy

migrate-up: ## Menjalankan database migrations
	@echo "Silakan jalankan migration SQL files di database/migrations/ secara manual atau gunakan tool migration"

migrate-down: ## Rollback database migrations
	@echo "Silakan rollback migration SQL files secara manual"

seed: ## Menjalankan database seeders
	@echo "Silakan jalankan seeder SQL files di database/seeders/ secara manual"

dev: ## Menjalankan dalam mode development dengan hot reload
	go run cmd/api/main.go

