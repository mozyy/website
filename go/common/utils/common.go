package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const configPath = "../../docker"

type Config struct {
	WebsocketPort string
	URL           string
	Websocket     string
	Database      struct {
		User     string
		Password string
		Domain   string
		Port     string
		Charset  string
	}
	DatamanageURL string
}

func GetConfig() Config {
	configByte, err := ioutil.ReadFile(configPath + "/config/config.json")
	PanicErr(err)
	var config map[string]Config
	json.Unmarshal(configByte, &config)
	env := GetEnv()
	return config[env]
}

func GetEnv() string {
	env := os.Getenv("GO_RUN_ENV")
	if env == "" {
		env = "production"
	}
	return env
}
