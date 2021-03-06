package command

import "net"

// CmdSet set命令
const CmdSet = "set"

func init() {
	Register(CmdSet, Set)
}

// Set set命令
func Set(conn net.Conn, args ...[]byte) {
	if len(args) < 2 {
		WriteError(conn, ErrWrongNumberArg.Errorf(CmdSet))
	} else {
		key := string(args[0])
		value := args[1]
		localNode.Set(key, value)
		WriteOK(conn)
	}

}
