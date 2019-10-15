package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"yyue.dev/common/message"
	"yyue.dev/common/utils"
	"yyue.dev/datamanage/database"
)

func main() {
	datamanageURL := utils.GetConfig().DatamanageURL
	q, err := database.New()
	utils.PanicErr(err)
	db, table := "development", "test"
	// q.DB.Exec("CREATE DATABASE " + db)
	err = q.ConnectDatabase(&db, &message.Message{})
	// _, err = q.DB.Exec("CREATE TABLE test (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(50), tel VARCHAR(50), PRIMARY KEY(id));")
	utils.PanicErr(err)
	err = q.ConnectTable(&table, &message.Message{})
	utils.PanicErr(err)
	err = rpc.Register(q)

	utils.PanicErr(err)

	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", datamanageURL)

	utils.PanicErr(err)
	log.Println("datamanage server started")
	err = http.Serve(listen, nil)
	utils.PanicErr(err)
}
