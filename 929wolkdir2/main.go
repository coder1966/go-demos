package main

import (
	"flag"
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"time"
)

// func visit(path string, di fs.DirEntry, err error) error {
// 	fmt.Printf("Visited: %s\n", path)
// 	return nil
// }

func main() {
	// kvs = kvs.Add("file_size",  filesize   , false, true)
	// kvs = kvs.Add("file_count", filecount , false, true)
	// kvs = kvs.Add("dir_count",  dircount , false, true)
	var (
		filesize  int64
		filecount int64
		dircount  int64
	)
	start := time.Now()
	flag.Parse()
	root := flag.Arg(0)
	if root == "" {
		root = "/"
	}
	// root := "/"
	// root = "/home/zhangub/testing"
	// root = "/home/zhangub"
	err := filepath.WalkDir(root, func(path string, di fs.DirEntry, err error) error {
		info, err := di.Info()
		if err != nil {
			fmt.Println(" error: ", err)
			return err
		}

		if info.IsDir() {
			dircount++
		} else {
			filecount++

		}
		filesize += info.Size()

		// fmt.Println(path, info.Size())
		return nil
	})
	fmt.Printf("filepath.WalkDir() returned %v\n", err)
	fmt.Println(filesize, filecount, dircount)
	fmt.Println("FINISH use time:", time.Since(start))
}

func isreg(filename string, regslice []string) bool {
	buf := filename
	flag := false
	for i := 0; i < len(regslice); i++ {
		reg := regexp.MustCompile(`^.+\.` + regslice[i] + `$`)
		result := reg.FindAllStringSubmatch(buf, 1)
		if len(result) != 0 {
			flag = true
			break
		}
	}
	return flag
}
