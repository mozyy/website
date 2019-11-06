package main

import (
	"log"
	"net/http"

	"go.yyue.dev/common/utils"
	"go.yyue.dev/white/websocket"
)

const configPath = "../../docker"

func main() {
	http.HandleFunc("/", websocket.Handler)
	// err := http.ListenAndServe(":6503", nil)
	config := utils.GetConfig()
	env := utils.GetEnv()
	var err error
	switch env {
	case "development":
		err = http.ListenAndServe(":"+config.WebsocketPort, nil)
	case "production":
		err = http.ListenAndServeTLS(":"+config.WebsocketPort, configPath+"/nginx/ssl/yyue.dev.crt", configPath+"/nginx/ssl/yyue.dev.key", nil)
	default:
		log.Println("unknow GO_RUN_ENV")
	}
	utils.PanicErr(err)
}
