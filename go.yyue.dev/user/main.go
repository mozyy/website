package main

import (
	"net/http"
	"net/rpc"

	"go.yyue.dev/common/utils"
	"go.yyue.dev/user/model"
)

func main() {

	user := new(model.User)
	err := rpc.Register(user)

	utils.PanicErr(err)

	rpc.HandleHTTP()

	err = http.ListenAndServe(":5000", nil)
	utils.PanicErr(err)

}
