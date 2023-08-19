package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	URL  string   `toml:"url,omitempty"` // Deprecated
	URLs []string `toml:"urls" json:"urls"`
}

func main() {

	s1 := Student{"abcd", []string{"xyz111", "xyz222"}}
	var s, _ = json.Marshal(s1)
	jsonStr := string(s)
	fmt.Println(jsonStr)
}
