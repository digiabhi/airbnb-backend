MIGRATIONS_FOLDER=db/migrations
DB_URL=root:root@tcp(localhost:3306)/auth_db

#Create a new migration -> make migrate-create name="your_migration_name"
migrate-create:
	goose -dir $(MIGRATIONS_FOLDER) create $(name) sql

migrate-up: #make migrate-up
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" up

migrate-down: #make migrate-down
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" down

# Rollback all migrations and reset database # make migrate-reset
migrate-reset:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" reset

# Show current migration status # make migrate-status
migrate-status:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" status

# Redo last migration (Down then Up) # make migrate-redo
migrate-redo:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" redo

# Run specific migration version # make migrate-version version=20200101120000
migrate-to:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" up-to $(version)

# Rollback to a specific migration version # make migrate-down-to version=20200101120000
migrate-down-to:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" down-to $(version)

# Force a specific migration version # make migrate-force version=20200101120000
migrate-force:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" force $(version)

# Print Goose help # make migrate-help
migrate-help:
	goose -h