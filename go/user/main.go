package main

import (
	"net/http"
	"net/rpc"

	"github.com/mozyy/utils"
)

func main() {

	user := User{}
	rpc.Register(user)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":5000", nil)
	utils.PanicErr(err)

}
