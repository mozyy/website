package database

import (
	"encoding/gob"
)

type DataRegister struct {
}

func Register(registers ...interface{}) {
	for _, reg := range registers {
		gob.Register(reg)
		// TODO: add automatic initialization database create tabales.
	}

}
