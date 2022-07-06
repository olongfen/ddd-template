package api

import (
	"embed"
	"github.com/swaggo/swag"
)

//go:embed generated_doc.swagger.json
var f embed.FS

var SwaggerInfo *swag.Spec

func init() {
	docTemplate, _ := f.ReadFile("generated_doc.swagger.json")
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
