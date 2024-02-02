package utils

import (
	"encoding/json"
)

func Any2Map(data any) (ans map[string]interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &ans)
	if err != nil {
		return
	}
	return
}
