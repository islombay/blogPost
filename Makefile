rest:
	go run cmd/main.go

grpc:
	go run cmd/main_grpc.go

docker-build:
	docker build -t blogpost .