package database

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	tests := map[string]interface{}{
		"":            time.Date(2019, 10, 1, 0, 0, 0, 0, time.UTC),
		"test string": "test string",
		"5":           5,
		"5.5":         5.5,
		"123123":      []interface{}{"a", "b", 1, time.Date(2019, 10, 1, 0, 0, 0, 0, time.UTC)},
		"point":       &([]string{"po", "i", "nt"}),
	}
	for want, value := range tests {
		got := format(reflect.ValueOf(value))
		if want != got {
			t.Errorf("format Time error, want: %s, got: %s", want, got)
		}
	}
}

func TestGetStructKV(t *testing.T) {
	tests := map[string]interface{}{
		"table": struct {
			Table string `db:"table"`
			id    string
		}{"table123", "1"},
		"one,two": struct {
			One string `db:"one"`
			Two string `db:"two"`
		}{"1", ""},
		"a": struct {
			A string `db:"a"`
			B string
		}{"a", "b"},
	}
	for want, value := range tests {

		got, v := getStructKV(value)
		if want != strings.Join(got, ",") {
			t.Errorf("get struct keys error, want: %s, got: %s, value: %s", want, strings.Join(got, ","), v)
		}
	}
}
