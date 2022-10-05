up:
	@echo "Starting Docker images..."
	docker compose up -d
down:
	@echo "Stopping Docker compose..."
	docker compose down
migrate_up:
	@echo "Running Migrations..."
	docker compose run --rm migrate up
migrate_down:
	@echo "Rollback Migrations..."
	docker compose run --rm migrate down
test:
	@echo "Running unit tests..."
	go test ./... -v -cover
