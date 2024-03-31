package utils

import (
	"encoding/json"
	"fmt"
)

func DisplayJSONOutput(stats interface{}) error {
	response, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(response))

	return nil
}
