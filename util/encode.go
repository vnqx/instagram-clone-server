package util

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func EncodeJSONBody(data interface{}) *bytes.Buffer {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(data); err != nil {
		fmt.Errorf("an error during encoding occurred: %s", err)
		panic(err)
	}
	return buf
}