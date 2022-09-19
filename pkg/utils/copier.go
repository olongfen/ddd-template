package utils

import jsoniter "github.com/json-iterator/go"

func Copier(in, out interface{}) (err error) {
	var (
		b []byte
	)

	if b, err = jsoniter.Marshal(in); err != nil {
		return
	}
	return jsoniter.Unmarshal(b, out)
}
