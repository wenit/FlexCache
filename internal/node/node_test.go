package node

import (
	"testing"
)

func TestNode_Get(t *testing.T) {
	node := NewLocalNode()
	key := "name"
	value := "abc"

	node.Set(key, []byte(value))
	data, _ := node.Get(key)
	t.Log(string(data))
}
