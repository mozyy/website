package database

import (
	"reflect"
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
