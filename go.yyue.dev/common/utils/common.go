package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
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
		env = "development"
	}
	return env
}

// WaitForIt is dependes for some tcp
func WaitForIt(tcp string) {
	if tcp == "" {
		return
	}
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		_, err := net.Dial("tcp", tcp)
		fmt.Println("err: ", err)
		if err == nil {
			return
		}
		<-ticker.C
	}
}
