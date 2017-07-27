package util

import (
	"encoding/json"
	"os"
)

func PrintPretty(v interface{}) (err error) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return
	}
	data = append(data, '\n')
	os.Stdout.Write(data)
	return
}
