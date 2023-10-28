run:
	nodemon --exec go run main.go --signal SIGTERM

check_install_swagger:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install_swagger
	GO111MODULE=off && swagger generate spec -o ./swagger.yaml --scan-models

postgres:
	docker run --name pg -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:latest

createdb:
	docker exec -it pg createdb --username=postgres --owner=postgres products

dropdb:
	docker exec -it pg dropdb --username=postgres products

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/products?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/products?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc