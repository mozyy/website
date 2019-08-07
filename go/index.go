package main

import (
	"net/http"

	"go/utils"
	"go/websocket"
)

func main() {
	http.HandleFunc("/", websocket.Handler)
	// err := http.ListenAndServe(":6503", nil)
	err := http.ListenAndServeTLS(":6503", "./config/other/fa19001a2e8fb15a.pem", "./config/other/fa19001a2e8fb15a.pem", nil)
	utils.PanicErr(err)
}
