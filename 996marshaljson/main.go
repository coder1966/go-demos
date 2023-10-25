package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	file := "data.json"

	f, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("read fail", err)
	}
	strs := strings.Split(string(f), "\n")

	res := []string{}

	for _, str := range strs {
		reKey := getWord(str)
		if reKey != "" {
			res = append(res, reKey)
		}
	}
	sort.Strings(strs)

	lines := `// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package 

import "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/plugins/inputs"

//nolint:lll
func (i *Input) Dashboard(lang inputs.I18n) map[string]string {
	switch lang {
	case inputs.I18nZh:
		return map[string]string{`

	fmt.Println(lines)

	for _, key := range res {
		fmt.Printf("\"%s\":\"\",\n", key)
	}

	lines = `}
	case inputs.I18nEn:
		return map[string]string{`

	fmt.Println(lines)

	for _, key := range res {
		fmt.Printf("\"%s\":\"\",\n", key)
	}

	lines = `}
	default:
		return nil
	}
}

func (i *Input) DashboardList() []string {
	return nil
}
`

	fmt.Println(lines)

}

func getWord(str string) string {
	s1 := strings.Split(str, "\"{{.Dashboard.")

	if len(s1) == 1 {
		return ""
	}

	s2 := strings.Split(s1[1], "}}\"")
	if len(s2) == 1 {
		return ""
	}

	return s2[0]

}
