package command

import (
	"aedis/internal/protocol"
	"net"
)

// CmdDel del
const CmdDel = "del"

func init() {
	Register(CmdDel, Del)
}

// Del 删除缓存数据
func Del(conn net.Conn, args ...[]byte) {
	if len(args) < 1 {
		WriteError(conn, ErrWrongNumberArg.Errorf(CmdDel))
	} else {
		count := 0
		for _, arg := range args {
			value, _ := localNode.Get(string(arg))
			if value != nil {
				localNode.Del(string(arg))
				count++
			}
		}
		conn.Write(protocol.MarshalIntergerReply(count))
	}

}
