package main

import (
	"net/http"

	"go/utils"
	"go/websocket"
)

func main() {
	http.HandleFunc("/", websocket.Handler)
	err := http.ListenAndServe(":6503", nil)
	// err := http.ListenAndServeTLS(":6503", "../docker/nginx/ssl/yyue.dev.crt", "../docker/nginx/ssl/yyue.dev.key", nil)
	utils.PanicErr(err)
}
