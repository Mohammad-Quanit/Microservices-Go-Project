git:
	git add .
	git commit -m "$(MSG)"
	git push

postgres:
	docker run --name pg -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:latest