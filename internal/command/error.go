package command

import (
	"aedis/internal/protocol"
	"fmt"
	"net"
)

// WriteError 返回错误回复
func WriteError(conn net.Conn, errorInfo interface{}) {
	err := fmt.Sprintf("%s%s%s", protocol.ReplyError, errorInfo, protocol.NewLine)
	conn.Write([]byte(err))
}
