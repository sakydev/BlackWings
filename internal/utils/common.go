package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

func DisplayJSONOutput(stats interface{}) error {
	response, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(response))

	return nil
}

func ImplodeString(a []string, separator string) string {
	return strings.Join(a, separator)
}

func ImplodeInt64(a []int64, separator string) string {
	var s []string
	for _, i := range a {
		s = append(s, fmt.Sprintf("%d", i))
	}

	return strings.Join(s, separator)
}
