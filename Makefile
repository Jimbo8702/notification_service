BINARY_NAME=myApp

build: 
	@go build -o bin/server .

run: build
	@./bin/server

test:
	@go test ./...

start: run

stop:
	@-pkill -SIGTERM -f "./bin/server"

proto: 
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative types/ptypes.proto

restart: stop start