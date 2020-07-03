test:
	go test ./... -cover

install:
	go mod tidy
	go mod download
	go install .

docker-build:
	docker build -t bookstore-api .

docker-run:
	docker run -h 127.0.0.1 -it --network=host -p 80:8081 bookstore-api
