package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"go.yyue.dev/common/types"
	"go.yyue.dev/common/utils"
	"go.yyue.dev/datamanage/database"
)

func main() {

	// register possible struct types for gob
	database.Register(types.HouseInfo{})

	datamanageURL := utils.GetConfig().DatamanageURL
	q, err := database.New()
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
