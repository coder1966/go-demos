package main

/*
	输入路径，
	读取这个路径的全部.go,md文件，放到切片，文件名，子目录_子目录_..._文件.go.html。
	注意win和linux的子目录组合方式。
	在目录下，创建 noteshtml 子目录，创建index.html，切片放进去成为索引，跳转的是相对路径
	读取.go .md go的只要有//就全行采用； 或分析/ * * /对，放到html，每行后面加上<br/>；md全体采用。
*/

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 文件列表
var list []string
var writePath string

func main() {
	list = make([]string, 0)
	writePath = filepath.Join(".", "")

	var path string
	fmt.Println("输入路径名称：")
	fmt.Scanln(&path)
	// 确保path目标存在
	isExists, _, err := IsPathExists()
	if err != nil {
		fmt.Println(" error: ", err)
		return
	}
	if !isExists {
		fmt.Println(" have not path: ", path)
		return
	}
	// 重新创建 noteshtml 子目录
	writePath, err = recreatPath(filepath.Join(writePath, "noteshtml"))

	fmt.Println("main 结束")
}

func traverseTree(path string) {
	fileInf, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(" error: ", err)
		return
	}
}

// IsPathExists Check if the given path exists and if is folder.
func IsPathExists(path string) (isExists, isDir bool, err error) {
	s, err := os.Stat(path)
	if err == nil {
		return true, s.IsDir(), nil
	}

	if os.IsNotExist(err) {
		return false, false, nil
	}

	return false, false, fmt.Errorf("check %v: %w", path, err)
}

func recreatPath(path string) (string, error) {
	s, err := os.Stat(path)

	if err == nil {
		// 存在
		err = os.RemoveAll(path)
		if err != nil {
			fmt.Println(" error: ", err)
			return "", err
		}
	} else if os.IsNotExist(err) {
		// 不存在
	} else {
		// 未知错误
		fmt.Println(" error: ", err)
		return "", err
	}

	// 创建

	return path, nil
}
