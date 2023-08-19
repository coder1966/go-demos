package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	file := "infile"

	f, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("read fail", err)
	}
	strs := strings.Split(string(f), "\n")

	sort.Strings(strs)

	for _, v := range strs {
		fmt.Println(v)

	}

}
