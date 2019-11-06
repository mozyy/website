package main

import (
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("start main")
	m.Run()
	log.Println("stop main")

}
