run:
	nodemon --exec go run *.go --signal SIGTERM

build:
	go build

check_install_swagger:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install_swagger
	GO111MODULE=off && swagger generate spec -o ./swagger.yaml --scan-models

startdb:
	docker start pg

postgres:
	docker run -d -e POSTGRES_PASSWORD=postgres --name pg -p 5432:5432 postgres

postgres-mquanit:
	docker run --name pg14 -p 5432:5432 -e POSTGRES_USER=quanit -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

migrateup:
	migrate -path db/migration -database "postgresql://quanit:secret@localhost:5432/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://quanit:secret@localhost:5432/simplebank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres postgres-mquanit migrateup migratedown sqlc