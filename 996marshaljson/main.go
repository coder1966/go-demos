package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var outFile = "out.go"
var inFile = "data.json"
var outLine string

func main() {

	f, err := os.ReadFile(inFile)
	if err != nil {
		panic(fmt.Sprintln("read fail", err))
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

package main

import "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/plugins/inputs"

//nolint:lll
func (i *Input) Dashboard(lang inputs.I18n) map[string]string {
	switch lang {
	case inputs.I18nZh:
		return map[string]string{`

	outLine += fmt.Sprintln(lines)

	for _, key := range res {
		tran := getTranslate(key, true)
		outLine += fmt.Sprintf("\"%s\":\"%s\",\n", key, tran)
	}

	lines = `}
	case inputs.I18nEn:
		return map[string]string{`

	outLine += fmt.Sprintln(lines)

	for _, key := range res {
		tran := getTranslate(key, false)
		outLine += fmt.Sprintf("\"%s\":\"%s\",\n", key, tran)
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

	outLine += fmt.Sprintln(lines)

	if err := os.WriteFile(outFile, []byte(outLine), 0666); err != nil {
		panic(err)
	}

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

func getTranslate(key string, isZh bool) string {

	ss, ok := newMap[key]
	if ok {
		if isZh {
			return ss[1]
		} else {
			return ss[0]
		}
	}

	return "################"
	panic(fmt.Sprint("not this key", key))
}
