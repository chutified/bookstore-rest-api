test:
	go test ./... -cover
build:
	docker build -t bookstore-api .
run:
	docker run -d -p 80:8081 bookstore-api
