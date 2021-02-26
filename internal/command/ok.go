package command

import (
	"aedis/internal/protocol"
	"net"
)

// WriteOK 返回状态回复
func WriteOK(conn net.Conn, args ...[]byte) {
	resp := WirteSingleLineReply(protocol.ReplyStatus, []byte(ReplyStatusOK))
	conn.Write(resp)
}
