package main

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

// Example of how to:
// 如何执行的示例：
// - Clone a repository into memory
// - 将存储库克隆到内存中
// - Get the HEAD reference
// - 获取 HEAD 参考
// - Using the HEAD reference, obtain the commit this reference is pointing to
// - 使用 HEAD 引用，获取该引用指向的提交
// - Using the commit, obtain its history and print it
// - 使用提交，获取其历史记录并打印它
func main() {
	// Clones the given repository, creating the remote, the local branches and fetching the objects, everything in memory:
	// 克隆给定的存储库，创建远程、本地分支并获取对象，内存中的所有内容：
	Info("git clone https://github.com/src-d/go-siva") // github.com/src-d/go-siva”）
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/src-d/go-siva", // github.com/src-d/go-siva",
	})
	CheckIfError(err)

	// Gets the HEAD history from HEAD, just like this command:
	// 从 HEAD 获取 HEAD 历史记录，就像这个命令：
	Info("git log")

	// ... retrieves the branch pointed by HEAD
	// ...检索 HEAD 指向的分支
	ref, err := r.Head()
	CheckIfError(err)

	// ... retrieves the commit history
	// ...检索提交历史记录
	since := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	until := time.Date(2019, 7, 30, 0, 0, 0, 0, time.UTC)
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: &since, Until: &until})
	CheckIfError(err)

	// ... just iterates over the commits, printing it
	// ...只是迭代提交，打印它
	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)

		return nil
	})
	CheckIfError(err)
}
