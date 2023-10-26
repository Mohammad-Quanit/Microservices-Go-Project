run:
	nodemon --exec go run main.go --signal SIGTERM

check_install_swagger:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install_swagger
	GO111MODULE=off && swagger generate spec -o ./swagger.yaml --scan-models

postgres-mquanit:
	docker run --name pg14 -p 5432:5432 -e POSTGRES_USER=quanit -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it pg14 createdb --username=quanit --owner=quanit simplebank

dropdb:
	docker exec -it pg14 dropdb simplebank

migrateup:
	migrate -path db/migration -database "postgresql://postgres@localhost:5432/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://quanit:secret@localhost:5432/simplebank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc