package message

import (
	"errors"
	"strconv"
)

type Status int

const (
	StateSuccess Status = iota
	StateInfo
	StateWarning
	StateError
)

func (s Status) String() string {
	if int(s) < len(statusNames) {
		return statusNames[s]
	}
	return "status" + strconv.Itoa(int(s))
}

var statusNames = []string{
	StateSuccess: "success",
	StateInfo:    "info",
	StateWarning: "warning",
	StateError:   "error",
}

type Message struct {
	Data    interface{}
	State   Status
	Message string
}

func (msg *Message) SetData(data interface{}) {
	msg.Data = data
}

// func (msg *Message) SetMessage(message string) {
// 	msg.Message = message
// }

func (msg *Message) Success(message string) {
	msg.State = StateSuccess
	msg.Message = message
}

func (msg *Message) Info(message string) {
	msg.State = StateInfo
	msg.Message = message
}

func (msg *Message) Warning(message string) {
	msg.State = StateWarning
	msg.Message = message
}

func (msg *Message) Error(message string) error {
	msg.State = StateError
	msg.Message = message
	return errors.New(message)
}

func Success(message string) Message {
	return Message{
		State:   StateSuccess,
		Message: message,
	}
}

func Info(message string) Message {
	return Message{
		State:   StateInfo,
		Message: message,
	}
}

func Warning(message string) Message {
	return Message{
		State:   StateWarning,
		Message: message,
	}
}

func Error(message string) Message {
	return Message{
		State:   StateError,
		Message: message,
	}
}
