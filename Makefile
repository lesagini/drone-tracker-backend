
createpostgrescontainer:
	docker run --name lima_drones -e POSTGRES_PASSWORD=secret -p 5000:5432 -d postgres:15.2-alpine
createdb:
	docker exec -it lima_drones createdb --username=postgres --owner=postgres drones

dropdb:
	docker exec -it lima_drones dropdb -U postgres drones

migrateup:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5000/drones?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5000/drones?sslmode=disable" -verbose down

.PHONY: