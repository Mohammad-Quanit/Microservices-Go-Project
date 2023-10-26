# A simple authentication & authorization service built on Go, Gin Framework, Go-JWT & Postgresql


Steps to Setup this workflow in your machine

1. Install Docker & Code Editor

2. Clone this repo by typing `git clone https://github.com/mohammad-quanit/auth-service.git` and go to auth-service directory by `cd auth-service`.

3. Run Command `go install && go mod tidy`

4. Run PostgreSQL using docker by runnning command `docker run -d -e POSTGRES_PASSWORD=postgres --name pg -p 5432:5432 postgres`

5. After running postgresql db run `go run main.go` and this will start go server.

6. Test the endpoints by using API testing tool like Postman or Thunder Client on `localhost:8080/signup` to create first user.