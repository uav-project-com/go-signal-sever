package lib

import (
	"encoding/json"
	"log"
)

func ToJsonStr(v interface{}) string {
	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Println("Error marshaling object:", err)
		return ""
	}
	return string(bytes)
}
