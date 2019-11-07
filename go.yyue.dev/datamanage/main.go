package main

import (
	"fmt"

	"github.com/micro/go-micro"
	"go.yyue.dev/datamanage/database"
	"go.yyue.dev/datamanage/proto"
)

func main() {
	srv := micro.NewService(
		micro.Name("database"),
	)
	srv.Init()
	proto.RegisterDatabaseServiceHandler(srv.Server(), database.New())
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
