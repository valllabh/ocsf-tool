package commons

import (
	"encoding/json"
	"fmt"
)

// func to PrintJson is a helper function to print a JSON object
// in a pretty format.
func PrintJson(obj any) {
	json, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json))
}
