package push

import (
	"encoding/json"
	"fmt"
	"log"

	webpush "github.com/SherClockHolmes/webpush-go"
	"yyue.dev/datamanage"
)

const (
	vapidPublicKey  = "BPb50zeEu3sM2AkuUlF0RwB5vN5_xWNI6FKHAacnmQqdVw3uVbIszG3yaGuR3UG566qvieQljTpINEXnEb3IghI"
	vapidPrivateKey = "rtjCOqWkLBaZorIoUixVI74WdiQ9wo62Ugk"
)

func getKey() {
	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		log.Fatal(err)
	}
	datamanage.GetDb("table")
	// utils.PanicErr(err)
	fmt.Printf("privateKey:\n%s\npublicKey\n%s\n", privateKey, publicKey)
}

//Get is get webpush subscription
func Get(subscription string) (*webpush.Subscription, error) {
	s := &webpush.Subscription{}
	err := json.Unmarshal([]byte(subscription), s)
	return s, err
}

// Send is send some message to web client
func Send(msg string, s *webpush.Subscription) error {
	_, err := webpush.SendNotification([]byte(msg), s, &webpush.Options{
		Subscriber:      "ex",
		VAPIDPublicKey:  vapidPublicKey,
		VAPIDPrivateKey: vapidPrivateKey,
	})
	return err
}
