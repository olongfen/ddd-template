package v1

import (
	"embed"
	"github.com/swaggo/swag"
)

//go:embed v1.swagger.json
var f embed.FS

var SwaggerInfo *swag.Spec

func init() {
	docTemplate, _ := f.ReadFile("v1.swagger.json")
	SwaggerInfo = &swag.Spec{
		Version:          "",
		Host:             "",
		BasePath:         "",
		Schemes:          []string{},
		Title:            "",
		Description:      "",
		InfoInstanceName: "swagger",
		SwaggerTemplate:  string(docTemplate),
	}
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)

}
