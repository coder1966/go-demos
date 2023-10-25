package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
)

// Basic example of how to clone a repository using clone options.
// 如何使用克隆选项克隆存储库的基本示例。
func main() {
	CheckArgs("<url>", "<directory>")
	url := os.Args[1]
	directory := os.Args[2]

	// Clone the given repository to the given directory
	// 将给定存储库克隆到给定目录
	Info("git clone %s %s --recursive", url, directory)

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	CheckIfError(err)

	// ... retrieving the branch being pointed by HEAD
	// ...检索 HEAD 指向的分支
	ref, err := r.Head()
	CheckIfError(err)
	// ... retrieving the commit object
	// ...检索提交对象
	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	fmt.Println(commit)
}
