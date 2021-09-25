package example3

import (
	"encoding/json"
)

func JSONToMap(b []byte) (map[string]interface{}, error) {
	var res map[string]interface{}
	err := json.Unmarshal(b, &res)

	return res, err
}
