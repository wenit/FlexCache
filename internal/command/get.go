package command

import (
	"aedis/internal/protocol"
	"net"
)

// CmdGet get命令
const CmdGet = "get"

func init() {
	Register(CmdGet, Get)
}

// Get get命令
func Get(conn net.Conn, args ...[]byte) {
	key := string(args[0])
	value, _ := localNode.Get(key)
	resp, _ := protocol.MarshalBytesArray(value)
	conn.Write(resp)
}
