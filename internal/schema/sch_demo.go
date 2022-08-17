package schema

type DemoResp struct {
	// 创建时间
	CreatedAt int64 `json:"createdAt"`
	// 消息
	Message string `json:"message"`
}
