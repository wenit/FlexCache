package cluster

import (
	"aedis/internal/node"
	"fmt"
	"strconv"
	"testing"
)

func TestCluster_Get(t *testing.T) {
	clusterSize := 3
	nodes := make([]*node.Node, 0)
	for i := 1; i <= clusterSize; i++ {
		ip := "192.168.0." + strconv.Itoa(i)
		node := node.NewNode(ip, 6379)
		fmt.Println(node)
		nodes = append(nodes, node)
	}
	cluster := CreateCluster(nodes)
	fmt.Println(cluster)

	counters := make(map[string]int)
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key-%d", i)
		solt := CalcSlot([]byte(key))
		fmt.Println(key, solt)
		node := cluster.GetNode(key)

		node.Set(key, []byte(key))

		id := node.ID
		if _, ok := counters[id]; !ok {
			counters[id] = 1
		} else {
			counters[id]++
		}
	}
	fmt.Println(counters)

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key-%d", i)
		v, _ := cluster.Get(key)
		t.Log(key, string(v))
	}
}
