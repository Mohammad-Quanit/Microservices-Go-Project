# A Microservice based application using Go

Currently there are 2 services right now, `Auth` and `Products`. This is a side project to practice microservices architecture and there is lot of other features coming. 

### Steps to up and run this project in your machine. (if you want)

1. Run PostgreSQL using docker by runnning command:
`docker run -d -e POSTGRES_PASSWORD=postgres --name pg -p 5432:5432 postgres`

<!-- docker run --name pg14 -e POSTGRES_USER=quanit -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:14-alpine
docker exec -it pg14 psql -U quanit -->