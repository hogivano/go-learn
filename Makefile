include .env

migrateup:
	migrate -path migrations -database "$(DB_MIGRATION)" -verbose up

migratedown:
	migrate -path migrations -database "$(DB_MIGRATION)" -verbose down

migratedrop:
	migrate -path migrations -database "$(DB_MIGRATION)" -verbose drop

migrateadd:
	migrate create -ext sql -dir migrations

.PHONY: migrateup migratedown migratedrop migrateadd