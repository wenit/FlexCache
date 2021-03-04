package command

import (
	"aedis/internal/protocol"
	"net"
)

// CmdExists 判断key是否存在
const CmdExists = "exists"

func init() {
	Register(CmdExists, Exists)
}

// Exists 判断key是否存在
func Exists(conn net.Conn, args ...[]byte) {
	key := string(args[0])
	value, _ := localNode.Get(key)
	if value == nil {
		resp := protocol.MarshalIntergerReply(0)
		conn.Write(resp)
	} else {
		resp := protocol.MarshalIntergerReply(1)
		conn.Write(resp)
	}
}
