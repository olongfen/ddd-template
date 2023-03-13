package mutation

// Mutation 操作变动的数据
type Mutation struct {
	demo IDemoService
}

// Demo 获取demo的写入服务
func (m *Mutation) Demo() IDemoService {
	return m.demo
}

// SetMutation set mutation
func SetMutation(demo IDemoService) *Mutation {
	return &Mutation{demo: demo}
}
