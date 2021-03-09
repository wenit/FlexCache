package node

import (
	"aedis/internal/slot"
	"encoding/json"
	"sync"

	"github.com/google/uuid"
)

// Node 节点
type Node struct {
	ID      string            `json:"id"`
	IP      string            `json:"ip"`
	Port    int               `json:"port"`
	MinSlot slot.Slot         `json:"minSlot"`
	MaxSlot slot.Slot         `json:"maxSlot"`
	Data    map[string][]byte `json:"data"`
	lock    sync.Mutex
}

func (n *Node) String() string {
	data, _ := json.Marshal(n)
	// return fmt.Sprintf("%s:%d", n.IP, n.Port)
	return string(data)
}

// Get 获取数据
func (n *Node) Get(key string) ([]byte, error) {
	n.lock.Lock()
	defer n.lock.Unlock()
	data, ok := n.Data[key]
	if ok {
		return data, nil
	}
	return nil, nil
}

// Set 设置数据
func (n *Node) Set(key string, value []byte) error {
	n.lock.Lock()
	defer n.lock.Unlock()
	n.Data[key] = value
	return nil
}

// Del 删除数据
func (n *Node) Del(key string) error {
	n.lock.Lock()
	defer n.lock.Unlock()
	delete(n.Data, key)
	return nil
}

// NewNode create node instance
func NewNode(ip string, port int) *Node {
	return &Node{
		ID:   uuid.NewString(),
		IP:   ip,
		Port: port,
		Data: make(map[string][]byte),
	}
}

// NewLocalNode 创建本地节点
func NewLocalNode() *Node {
	return &Node{
		ID:   uuid.NewString(),
		IP:   "127.0.0.1",
		Port: 6379,
		Data: make(map[string][]byte),
	}
}
