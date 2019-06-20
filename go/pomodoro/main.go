package main

import (
	"net/rpc"

	"github.com/mozyy/utils"
)

func main() {
	client, err := rpc.DialHTTP("tcp", ":5000")
	utils.PanicErr(err)
	defer client.Close()

	client.Call("User.add", nil, nil)
}
