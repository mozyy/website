package message

import "errors"

type Message struct {
	Data    interface{}
	State   string
	Message string
}

func (msg *Message) setData(data interface{}) {
	msg.Data = data
}

func (msg *Message) setMessage(message string) {
	msg.Message = message
}

func (msg *Message) Success(message string) {
	msg.State = "success"
	msg.Message = message
}

func (msg *Message) Info(message string) {
	msg.State = "info"
	msg.Message = message
}

func (msg *Message) Warning(message string) {
	msg.State = "warning"
	msg.Message = message
}

func (msg *Message) Error(message string) error {
	msg.State = "error"
	msg.Message = message
	return errors.New(message)
}

func Success(message string) Message {
	return Message{
		State:   "success",
		Message: message,
	}
}

func Info(message string) Message {
	return Message{
		State:   "info",
		Message: message,
	}
}

func Warning(message string) Message {
	return Message{
		State:   "warning",
		Message: message,
	}
}

func Error(message string) Message {
	return Message{
		State:   "error",
		Message: message,
	}
}
