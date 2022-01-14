# Go Microservices App

docker run --name pg14 -e POSTGRES_USER=quanit -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:14-alpine
docker exec -it pg14 psql -U quanit