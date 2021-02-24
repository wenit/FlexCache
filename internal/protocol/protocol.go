package protocol

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

// ErrProtocolEncodeEmpty empty cmd
var (
	ErrProtocolEncodeEmpty = errors.New("encode empty cmd")
)

const (
	// ReplyStatus 状态回复
	ReplyStatus = "+"
	// ReplyError 错误回复
	ReplyError = "-"
	// ReplyInteger 整数回复
	ReplyInteger = ":"
	// ReplyBulk 批量回复
	ReplyBulk = "$"
	// ReplyMultiReply 多条批量回复
	ReplyMultiReply = "*"
	// NewLine 换行
	NewLine = "\r\n"
	// Space 分隔符
	Space = " "
)

// MarshalStr 字符串编码
func MarshalStr(cmd string) ([]byte, error) {
	return MarshalBytes([]byte(cmd))
}

// MarshalBytes 字节数组编码
func MarshalBytes(cmd []byte) ([]byte, error) {
	r := bytes.Split(cmd, []byte(Space))
	if r == nil || len(r) == 0 {
		return nil, ErrProtocolEncodeEmpty
	}
	var buffer bytes.Buffer
	var reqBuffer bytes.Buffer
	argsLen := 0
	for _, v := range r {
		if len(v) > 0 {
			argsLen++
			argLen := len(v)
			buffer.WriteString(ReplyBulk + strconv.Itoa(argLen) + NewLine)
			buffer.Write(v)
			buffer.WriteString("\r\n")
		}
	}
	reqBuffer.WriteString(ReplyMultiReply + strconv.Itoa(argsLen) + NewLine)
	reqBuffer.Write(buffer.Bytes())
	return reqBuffer.Bytes(), nil
}

// UnMarshalBytes 字节数组解码
func UnMarshalBytes(cmd []byte) {
	r := bytes.Split(cmd, []byte(NewLine))
	for _, v := range r {
		fmt.Println(string(v))
	}

	// first := cmd[0]
	// if strings.EqualFold(ReplyMultiReply, string(first)) {

	// }

	// switch first {
	// 	case ReplyMultiReply
	// }
}
