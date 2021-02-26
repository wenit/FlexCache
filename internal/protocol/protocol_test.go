package protocol

import (
	"math"
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

func TestUnMarshalBytes(t *testing.T) {
	resps := []string{
		"+OK\r\n",
		"+PONG\r\n",
		"-Err unknown command\r\n",
	}
	for _, resp := range resps {

		vv := UnMarshalBytes([]byte(resp))
		for _, v := range vv {
			t.Log(string(v))
		}

	}
}

func TestIntToBytes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1<<0",
			args: args{1 << 0},
		},
		{
			name: "math.MaxInt8",
			args: args{math.MaxInt8},
		},
		{
			name: "math.MaxInt16",
			args: args{math.MaxInt16},
		},
		{
			name: "math.MaxInt32",
			args: args{math.MaxInt32},
		},
		{
			name: "math.MaxInt64",
			args: args{math.MaxInt64},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes := IntToBytes(tt.args.n)
			o, _ := BytesToInt(bytes)
			t.Log("==", o)
		})
	}
}

func BenchmarkIntToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntToBytes(math.MaxInt32)
	}
}
