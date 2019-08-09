package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go/utils"
	"go/websocket"
)

func main() {
	http.HandleFunc("/", websocket.Handler)
	// err := http.ListenAndServe(":6503", nil)
	config := getConfig()
	err := http.ListenAndServeTLS(":"+config["websocketPort"].(string), "../docker/nginx/ssl/yyue.dev.crt", "../docker/nginx/ssl/yyue.dev.key", nil)
	utils.PanicErr(err)
}

func getConfig() map[string]interface{} {
	configByte, err := ioutil.ReadFile("../docker/config/config.json")
	utils.PanicErr(err)
	var config map[string]interface{}
	json.Unmarshal(configByte, &config)
	return config
}
