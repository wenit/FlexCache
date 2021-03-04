package command

import (
	"aedis/internal/errors"
	"aedis/internal/node"
	"aedis/internal/protocol"
	"bytes"
	"net"
	"strings"
)

var (
	// ErrUnknownCommand 未知的命令
	ErrUnknownCommand = errors.New("ERR unknown command '%s'")
	// ErrWrongNumberArg 参数数量错误
	ErrWrongNumberArg = errors.New("ERR wrong number of arguments for '%s' command")
)

const (
	// ReplyStatusOK 状态回复OK
	ReplyStatusOK = "OK"
	// ReplyStatusPong 状态回复PONG
	ReplyStatusPong = "PONG"
)

var localNode = node.NewLocalNode()

// Command 命令行信息
type commandInfo struct {
	cmds map[string]func(conn net.Conn, args ...[]byte)
}

var defaultCommand commandInfo = commandInfo{
	cmds: make(map[string]func(conn net.Conn, args ...[]byte)),
}

func (c *commandInfo) put(name string, fun func(conn net.Conn, args ...[]byte)) {
	name = strings.ToLower(name)
	c.cmds[name] = fun
}

func (c *commandInfo) get(name string) func(conn net.Conn, args ...[]byte) {
	name = strings.ToLower(name)
	return c.cmds[name]
}

func (c *commandInfo) contains(name string) bool {
	name = strings.ToLower(name)
	_, ok := c.cmds[name]
	return ok
}

// Contains 是否包含命令
func Contains(name string) bool {
	return defaultCommand.contains(name)
}

// GetCommand 获取命令执行函数
func GetCommand(name string) func(conn net.Conn, args ...[]byte) {
	return defaultCommand.get(name)
}

// Register 注册函数
func Register(name string, fun func(conn net.Conn, args ...[]byte)) {
	defaultCommand.put(name, fun)
}

// WirteSingleLineReply 返回单独行
func WirteSingleLineReply(replyType string, data []byte) []byte {
	buf := make([]byte, 0)
	buffer := bytes.NewBuffer(buf)
	buffer.WriteString(replyType)
	buffer.Write(data)
	buffer.WriteString(protocol.NewLine)
	return buffer.Bytes()
}
