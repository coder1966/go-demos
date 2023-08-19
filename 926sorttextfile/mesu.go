package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func mesu() {
	file := "mesuStr"

	f, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("read fail", err)
	}
	strs := strings.Split(string(f), "\n")

	for _, v := range strs {
		if strings.Index(v, "inputs.Summary") > 0 {
			fmt.Println(v)
		}

	}

}
