migrateup:
	migrate -path migrations -database "postgresql://admin:admin@localhost:5432/instagram?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgresql://admin:admin@localhost:5432/instagram?sslmode=disable" -verbose down

.PHONY: migrateup migratedown