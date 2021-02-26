package command

import (
	"aedis/internal/node"
	"net"
	"strings"
)

var localNode = node.NewLocalNode()

// Command 命令行信息
type commandInfo struct {
	cmds map[string]func(conn net.Conn, args ...[]byte)
}

var defaultCommand commandInfo

func init() {
	defaultCommand = commandInfo{
		cmds: make(map[string]func(conn net.Conn, args ...[]byte)),
	}
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
