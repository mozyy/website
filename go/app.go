package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"website-go/utils"
	"website-go/websocket"
)

func main() {
	http.HandleFunc("/", websocket.Handler)
	// err := http.ListenAndServe(":6503", nil)
	config := getConfig()
	env := getEnv()
	var err error
	if env == "development" {
		err = http.ListenAndServe(":"+config["websocketPort"].(string), nil)
	} else if env == "production" {
		err = http.ListenAndServeTLS(":"+config["websocketPort"].(string), "../docker/nginx/ssl/yyue.dev.crt", "../docker/nginx/ssl/yyue.dev.key", nil)
	}
	switch env {
	case "development":
		err = http.ListenAndServe(":"+config["websocketPort"].(string), nil)
	case "production":
		err = http.ListenAndServeTLS(":"+config["websocketPort"].(string), "../docker/nginx/ssl/yyue.dev.crt", "../docker/nginx/ssl/yyue.dev.key", nil)
	default:
		log.Println("unknow GO_RUN_ENV")
	}
	utils.PanicErr(err)
}

func getConfig() map[string]interface{} {
	configByte, err := ioutil.ReadFile("../docker/config/config.json")
	utils.PanicErr(err)
	var config map[string]interface{}
	json.Unmarshal(configByte, &config)
	return config
}

func getEnv() string {
	env := os.Getenv("GO_RUN_ENV")
	if env == "" {
		env = "production"
	}
	return env
}
