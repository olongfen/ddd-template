package schema

import "ddd-template/internal/common/utils"

type DemoResp struct {
	// 创建时间
	CreatedAt utils.JSONTime `json:"createdAt" swaggertype:"integer"`
	// 消息
	Message string `json:"message"`
}
