package command

// Command 操作变动的数据
type Command struct {
	demo IDemoService
}

// Demo 获取demo的写入服务
func (m *Command) Demo() IDemoService {
	return m.demo
}

// SetMutation set mutation
func SetMutation(demo IDemoService) *Command {
	return &Command{demo: demo}
}
