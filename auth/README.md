# A simple authentication & authorization service built on Go, Gin Framework, Go-JWT & Postgresql


Steps to Up and run auth service in your machine

1. Go to auth directory by `cd auth`.

2. Run Command `go install && go mod tidy`

3. Run `make run` and this will start authentication server.

4. Test the endpoints by using API testing tool like Postman or Thunder Client on `localhost:8080/signup` to create first user.