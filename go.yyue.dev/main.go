package main

// 不能一次protoc 所有porot文件
//go:generate protoc --proto_path=./proto --go_out=plugins=micro:../ ./proto/datamanage/database.proto
//go:generate protoc --proto_path=./proto --go_out=plugins=micro:../ ./proto/crawler/lianjia.proto
//go:generate protoc --proto_path=./proto --go_out=plugins=micro:../ ./proto/common/message.proto

func main() {

}
