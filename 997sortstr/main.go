package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	str := ``
	if str == "" {
		f, err := ioutil.ReadFile("old.txt")
		if err != nil {
			panic("read fail")
		}
		str = string(f)
	}
	strs := strings.Split(str, "\n")
	sort.Slice(strs, func(i, j int) bool {
		return strs[i] < strs[j]
	})

	str = ""
	for _, v := range strs {
		str += fmt.Sprintln(v)

	}

	fmt.Println(str)
}
