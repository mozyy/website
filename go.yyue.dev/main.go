package main

// 不能一次protoc 所有porot文件
//go:generate protoc --proto_path=./proto --go_out=plugins=grpc:../ ./proto/database.proto
//go:generate protoc --proto_path=./proto --go_out=plugins=grpc:../ ./proto/houseinfo.proto
//go:generate protoc --proto_path=./proto --go_out=plugins=grpc:../ ./proto/message.proto

func main() {

}
