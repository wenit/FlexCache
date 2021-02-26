package command

import (
	"aedis/internal/protocol"
	"net"
)

// CmdPing ping命令
const CmdPing = "ping"

func init() {
	Register(CmdPing, Ping)
}

// Ping ping命令处理
func Ping(conn net.Conn, args ...[]byte) {
	resp := WirteSingleLineReply(protocol.ReplyStatus, []byte(ReplyStatusPong))
	conn.Write(resp)
}
