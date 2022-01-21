package commandHandler

import (
	"encoding/json"
	"fmt"
)

func PrintJson(v interface{}) {
	b, err := json.MarshalIndent(v, "", "    ")
	fmt.Println(err, string(b))
}
