package command

import (
	"aedis/internal/errors"
	"aedis/internal/protocol"
	"bytes"
)

var (
	// ErrUnknownCommand 未知的命令
	ErrUnknownCommand = errors.New("Err unknown command '%s'")
)

const (
	// ReplyStatusOK 状态回复OK
	ReplyStatusOK = "OK"
	// ReplyStatusPong 状态回复PONG
	ReplyStatusPong = "PONG"
)

// WirteSingleLineReply 返回单独行
func WirteSingleLineReply(replyType string, data []byte) []byte {
	buf := make([]byte, 0)
	buffer := bytes.NewBuffer(buf)
	buffer.WriteString(replyType)
	buffer.Write(data)
	buffer.WriteString(protocol.NewLine)
	return buffer.Bytes()
}
