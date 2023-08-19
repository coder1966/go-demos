package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f1 := "out_couchbase2"
	f2 := "out_prom2"
	key, _ := ioutil.ReadFile(f1)
	s1 := string(key)
	s1 = replace(s1)
	key, _ = ioutil.ReadFile(f2)
	s2 := string(key)
	s2 = replace(s2)
	ss1 := strings.Split(s1, "\n")
	ss2 := strings.Split(s2, "\n")
	ss1 = deleteLine(ss1)
	ss2 = deleteLine(ss2)
	sort.Strings(ss1)
	sort.Strings(ss2)
	for i := 0; i < len(ss1); i++ {
		ok, rate := compare(ss1[i], ss2[i])
		if !ok || rate != 0 {
			fmt.Println(i, ok, rate, ss1[i])
			fmt.Println(i, ok, rate, ss2[i])
		}
		// fmt.Println(i, ok, rate, ss1[i])

	}
}

func replace(s string) string {
	s = strings.ReplaceAll(s, ",host=zhangub-OMEN-by-HP-Laptop-15-dc1xxx", "")
	s = strings.ReplaceAll(s, ",instance=127.0.0.1:9091", "")
	s = strings.ReplaceAll(s, ",host=192.168.56.14", "")
	return s
}
func deleteLine(s []string) []string {
	for i := len(s) - 1; i >= 0; i-- {
		if strings.Contains(s[i], " up=") ||
			strings.Contains(s[i], " bucketstats_up=") ||
			strings.Contains(s[i], " bucketstats_scrape_duration_seconds=") ||
			strings.Contains(s[i], "go,") ||
			strings.Contains(s[i], "go ") ||
			strings.Contains(s[i], "process ") ||
			strings.Contains(s[i], "promhttp,") ||
			strings.Contains(s[i], "promhttp ") ||
			strings.Contains(s[i], " scrape_duration_seconds=") {
			s = append(s[:i], s[i+1:]...)
		}

	}
	return s
}

func compare(s1, s2 string) (bool, float64) {
	// 去掉尾巴时间戳
	// s1 = s1[:strings.LastIndex(s1, " ")]
	// s2 = s2[:strings.LastIndex(s2, " ")]

	if s1 == s2 {
		return true, 0
	}

	s1h := s1[:strings.LastIndex(s1, "=")]
	s2h := s2[:strings.LastIndex(s2, "=")]
	s1t := s1[strings.LastIndex(s1, "=")+1:]
	s2t := s2[strings.LastIndex(s2, "=")+1:]

	ok := s1h == s2h
	s1d, _ := strconv.ParseFloat(s1t, 64)
	s2d, _ := strconv.ParseFloat(s2t, 64)
	rate := float64(-1)
	if s2d > 0 {
		rate = float64(s1d) / float64(s2d)
	}

	return ok, rate

}
