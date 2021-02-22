package command

// Command 命令接口
type Command interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}
