package store

import (
	"bytes"
	"testing"

	"go.yyue.dev/user/proto"
)

func TestEditTemplate(t *testing.T) {
	buf := &bytes.Buffer{}
	args := []proto.UserInfo{
		{Name: "a"},
		{Name: ""},
	}
	for _, v := range args {
		buf.Reset()
		err := editTemplate.Execute(buf, v)
		t.Errorf("template: %s", buf.Bytes())
		if err != nil {
			t.Errorf("error: %s", err.Error())
		}
	}

}
