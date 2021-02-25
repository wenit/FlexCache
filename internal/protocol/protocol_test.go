package protocol

import (
	"testing"
)

func TestMarshalStr(t *testing.T) {
	cmds := []string{
		"set name abc",
		"get name",
		"set  name  abc",
		"set   name     abc   ",
		"set   name    \" abc  \" ",
	}
	for _, cmd := range cmds {
		marshal, _ := MarshalStr(cmd)
		t.Log(string(marshal))

		args := UnMarshalBytes(marshal)
		argStr := ""
		for _, arg := range args {
			argStr += string(arg) + " "
		}
		t.Log(argStr)
	}
}
