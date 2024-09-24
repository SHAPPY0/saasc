package utils

import (
	"encoding/json"
)

func Stringify(st interface{}) string {
	str, err := json.Marshal(st)
	if err != nil {
		return ""
	}
	return string(str)
}