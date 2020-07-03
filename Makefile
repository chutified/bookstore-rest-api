test:
	go test ./... -cover
build:
	docker build -t bookstore-api .
run:
	docker run -h 127.0.0.1 -it --network=host -p 80:8081 bookstore-api
