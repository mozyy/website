package message

import (
	"errors"
)

func (msg *Message) Success(message string) {
	msg.State = Status_Success
	msg.Message = message
}

func (msg *Message) Info(message string) {
	msg.State = Status_Info
	msg.Message = message
}

func (msg *Message) Warning(message string) {
	msg.State = Status_Warning
	msg.Message = message
}

func (msg *Message) Error(message string) error {
	msg.State = Status_Error
	msg.Message = message
	return errors.New(message)
}

func Success(message string) Message {
	return Message{
		State:   Status_Success,
		Message: message,
	}
}

func Info(message string) Message {
	return Message{
		State:   Status_Info,
		Message: message,
	}
}

func Warning(message string) Message {
	return Message{
		State:   Status_Warning,
		Message: message,
	}
}

func Error(message string) Message {
	return Message{
		State:   Status_Error,
		Message: message,
	}
}
