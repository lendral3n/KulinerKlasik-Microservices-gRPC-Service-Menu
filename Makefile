proto:
	protoc --proto_path=./pkg/pb --go_out=. --go-grpc_out=. ./pkg/pb/*.proto
server:
	go run cmd/main.go