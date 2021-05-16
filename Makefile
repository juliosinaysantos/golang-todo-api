.PHONY: dev
dev:
	go run cmd/main.go

.PHONY: migrateUp
migrateUp:
	dbmate up

.PHONY: migrateDown
migrateDown:
	dbmate down