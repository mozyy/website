package main

import (
	"net/http"

	"go/utils"
	"go/websocket"
)

func main() {
	http.HandleFunc("/", websocket.Handler)
	// err := http.ListenAndServe(":6503", nil)
	err := http.ListenAndServeTLS(":6503", "./config/gd_bundle-g2-g1.crt", "./config/generated-private-key.txt", nil)
	utils.PanicErr(err)
}
