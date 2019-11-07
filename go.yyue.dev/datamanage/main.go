package main

import (
	"github.com/micro/go-micro"
	"go.yyue.dev/datamanage/database"
	"go.yyue.dev/datamanage/proto"
)

func main() {
	srv := micro.NewService()
	srv.Init()
	proto.RegisterDatabaseServiceHandler(srv.Server(), database.New())
}
