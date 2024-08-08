package main

import (
	"encoding/json"
	"fmt"
)

func printJson(v any) error {
	res, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(res))
	return nil
}
