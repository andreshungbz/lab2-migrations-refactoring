-include .envrc

# GO RULES

# run: Runs the main.go file
.PHONY: run
run:
	@echo '[MAKE] Running program...'
	go run ./cmd/api -port=4000

# format: Runs go fmt on all files
.PHONY: format
format:
	@echo '[MAKE] Formatting code...'
	go fmt ./...

# vet: Runs go vet on all files
.PHONY: vet
vet:
	@echo '[MAKE] Vetting code...' 
	go vet ./...

# test: Runs all Go tests verbosely
.PHONY: test
test:
	@echo '[MAKE] Testing code...'
	go test -v ./...

# build: Formats, vets, tests code; then builds binary file
.PHONY: build
build: format vet test
	@echo '[MAKE] Building program binary...'
	go build -o go-webapp ./cmd/api

# clean: Deletes binary file
.PHONY: clean
clean:
	@echo '[MAKE] Deleting program binary...'
	rm -rf ./go-webapp

# DATABASE RULES

# db/psql: Connects to the database with the appropriate user
.PHONY: db/psql
db/psql:
	psql ${DB_DSN}

# db/migrations/new name=$1: Create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo '[MAKE] Creating golang-migrate migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

# db/migrations/up: Apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up:
	@echo '[MAKE] Running golang-migrate up migrations...'
	migrate -path=./migrations -database=${DB_DSN} up
