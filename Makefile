test:
	go test ./... -cover
run:
	docker build -t bookstore-api Dockerfile
	docker run -d -p 80:8081
