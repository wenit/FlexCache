package main

import (
	"aedis/internal/command"
	"aedis/internal/node"
	"aedis/internal/protocol"
	"bufio"
	"fmt"
	"net"
	"strings"
)

var localNode = node.NewLocalNode()

func main() {

	fmt.Println("start server...")

	listener, err := net.Listen("tcp", "127.0.0.1:6379")

	if err != nil {
		fmt.Println("error listening", err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting", err.Error())
			return
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		cmds := protocol.UnMarshal(reader)
		for _, cmd := range cmds {
			fmt.Println(string(cmd))
		}
		// conn.Write([]byte("aa"))

		handleCommand(conn, cmds...)
	}
}

func handleCommand(conn net.Conn, cmds ...[]byte) {
	cmd := cmds[0]
	cmdStr := string(cmd)
	cmdStr = strings.ToLower(cmdStr)
	if !command.Contains(cmdStr) {
		command.WriteError(conn, command.ErrUnknownCommand.Errorf(cmdStr))
		WriteError(conn, command.ErrUnknownCommand.Error())
	} else {
		fun := command.GetCommand(cmdStr)
		args := cmds[1:]
		fun(conn, args...)
	}
}

// WriteError 返回错误回复
func WriteError(conn net.Conn, errorInfo interface{}) {
	err := fmt.Sprintf("%s%s%s", protocol.ReplyError, errorInfo, protocol.NewLine)
	conn.Write([]byte(err))
}

// WriteOk 返回状态回复
func WriteOk(conn net.Conn) {
	err := fmt.Sprintf("%s%s%s", protocol.ReplyStatus, "OK", protocol.NewLine)
	conn.Write([]byte(err))
}
