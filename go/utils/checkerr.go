package utils

import "log"

// PanicErr is commom check error info
func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

// LogErr is commom check error info
func LogErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
