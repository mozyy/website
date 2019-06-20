package message

import "errors"

type Message struct {
	data    interface{}
	state   string
	message string
}

func (msg *Message) setData(data interface{}) {
	msg.data = data
}

func (msg *Message) setMessage(message string) {
	msg.message = message
}

func (msg *Message) Success(message string) {
	msg.state = "success"
	msg.message = message
}

func (msg *Message) Info(message string) {
	msg.state = "info"
	msg.message = message
}

func (msg *Message) Warning(message string) {
	msg.state = "warning"
	msg.message = message
}

func (msg *Message) Error(message string) error {
	msg.state = "error"
	msg.message = message
	return errors.New(message)
}

func Success(message string) Message {
	return Message{
		state:   "success",
		message: message,
	}
}

func Info(message string) Message {
	return Message{
		state:   "info",
		message: message,
	}
}

func Warning(message string) Message {
	return Message{
		state:   "warning",
		message: message,
	}
}

func Error(message string) Message {
	return Message{
		state:   "error",
		message: message,
	}
}
