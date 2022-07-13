package api

import (
	"embed"
	jsoniter "github.com/json-iterator/go"
	"github.com/swaggo/swag"
)

//go:embed generated_doc.swagger.json
var f embed.FS

//go:embed swagger.json
var f1 embed.FS

var SwaggerInfo *swag.Spec

func init() {
	var (
		docTemplMap     = make(map[string]interface{})
		customizeDocMap = make(map[string]interface{})
	)
	docTemplate, _ := f.ReadFile("generated_doc.swagger.json")
	customizeDoc, _ := f1.ReadFile("swagger.json")
	_ = jsoniter.Unmarshal(docTemplate, &docTemplMap)
	_ = jsoniter.Unmarshal(customizeDoc, &customizeDocMap)
	for k, v := range customizeDocMap["definitions"].(map[string]interface{}) {
		docTemplMap["definitions"].(map[string]interface{})[k] = v
	}

	for k, v := range customizeDocMap["paths"].(map[string]interface{}) {
		docTemplMap["paths"].(map[string]interface{})[k] = v
	}
	docTemplate, _ = jsoniter.Marshal(docTemplMap)
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
