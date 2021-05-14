postgres:
	docker run --name postgres-13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13.2-alpine

createdb:
	docker exec -it postgres-13 createdb --username=root --owner=root RoaminRoninXBank

dropdb: 
	docker exec -it postgres-13 dropdb RoaminRoninXBank

migrateup: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/RoaminRoninXBank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/RoaminRoninXBank?sslmode=disable" -verbose down

sqlcG:
	sqlc generate


test:
	go test -v -cover ./...
	
.PHONY: postgres createdb dropdb migrateup migratedown sqlcG test
