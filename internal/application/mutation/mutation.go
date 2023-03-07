package mutation

import "ddd-template/internal/application/mutation/mutat_iface"

// Mutation 操作变动的数据
type Mutation struct {
	demo mutat_iface.IDemoService
}

// Demo 获取demo的写入服务
func (m *Mutation) Demo() mutat_iface.IDemoService {
	return m.demo
}

// SetMutation set mutation
func SetMutation(demo mutat_iface.IDemoService) *Mutation {
	return &Mutation{demo: demo}
}
