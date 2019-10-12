package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const configPath = "../../docker"

func GetConfig() map[string]interface{} {
	configByte, err := ioutil.ReadFile(configPath + "/config/config.json")
	PanicErr(err)
	var config map[string]interface{}
	json.Unmarshal(configByte, &config)
	return config
}

func GetEnv() string {
	env := os.Getenv("GO_RUN_ENV")
	if env == "" {
		env = "production"
	}
	return env
}
