package node

import "time"

// HotKey 热点数据
type HotKey struct {
	Key          string    `json:"key"`          // key 名称
	Size         int       `json:"size"`         // 数据大小
	LastModified time.Time `json:"lastModified"` // 上次修改时间
	LastAccess   time.Time `json:"lastAccess"`   // 上次访问时间
}

// HotKeys 热点数据集
type HotKeys struct {
	hotKeys map[string]*HotKey
}

// Put 设置热点数据
func (h *HotKeys) Put(key string) {}

// Remove 移除热点数据
func (h *HotKeys) Remove(key string) {}
