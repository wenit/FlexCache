package cluster

import (
	"aedis/internal/node"
	"aedis/internal/slot"
	"aedis/pkg/crc16"

	"github.com/google/uuid"
)

// Cluster 集群
type Cluster struct {
	ID    string
	Nodes []*node.Node
}

// Get 获取数据
func (c *Cluster) Get(key string) ([]byte, error) {
	node := c.GetNode(key)
	return node.Get(key)
}

// Set 设置数据
func (c *Cluster) Set(key string, value []byte) error {
	node := c.GetNode(key)
	return node.Set(key, value)
}

// GetNode 计算key所在的节点
func (c *Cluster) GetNode(key string) *node.Node {
	slot := slot.Slot(CalcSlot([]byte(key)))
	nodes := c.Nodes
	for _, node := range nodes {
		if slot >= node.MinSlot && slot <= node.MaxSlot {
			return node
		}
	}
	return nil
}

// CreateCluster 创建集群
func CreateCluster(Nodes []*node.Node) *Cluster {
	cluster := &Cluster{
		ID:    uuid.NewString(),
		Nodes: Nodes,
	}
	cluster.allocSlot()
	return cluster
}

func (c *Cluster) allocSlot() {
	nodes := c.Nodes
	size := len(nodes)

	rangeSlots := autoAllocSlot(size)

	for i, node := range nodes {
		node.MinSlot = rangeSlots[i].MinSlot
		node.MaxSlot = rangeSlots[i].MaxSlot
	}
}

// CalcSlot cacl slot
func CalcSlot(key []byte) uint16 {
	sum := crc16.CheckSum(key)
	slot := sum & slot.MaxSlotSize
	return slot
}

func autoAllocSlot(size int) []*slot.RangeSlot {
	l := slot.MaxSlotSize / size
	rangeSlots := make([]*slot.RangeSlot, 0)
	for i := 0; i < size-1; i++ {
		minSlot := slot.Slot(i * l)
		maxSlot := slot.Slot((i+1)*l - 1)
		rangeSlot := &slot.RangeSlot{
			MinSlot: minSlot,
			MaxSlot: maxSlot,
		}
		rangeSlots = append(rangeSlots, rangeSlot)
	}
	minSlot := slot.Slot((size - 1) * l)
	maxSlot := slot.Slot(slot.MaxSlotSize)
	rangeSlot := &slot.RangeSlot{
		MinSlot: minSlot,
		MaxSlot: maxSlot,
	}
	rangeSlots = append(rangeSlots, rangeSlot)
	return rangeSlots
}
