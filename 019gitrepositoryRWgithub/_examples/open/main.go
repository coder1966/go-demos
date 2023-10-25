package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Open an existing repository in a specific folder.
// 打开特定文件夹中的现有存储库。
func main() {
	CheckArgs("<path>")
	path := os.Args[1]

	// We instantiate a new repository targeting the given path (the .git folder)
	// 我们实例化一个针对给定路径（.git 文件夹）的新存储库
	r, err := git.PlainOpen(path)
	CheckIfError(err)

	// Length of the HEAD history
	// HEAD 历史记录的长度
	Info("git rev-list HEAD --count")

	// ... retrieving the HEAD reference
	// ...检索 HEAD 引用
	ref, err := r.Head()
	CheckIfError(err)

	// ... retrieves the commit history
	// ...检索提交历史记录
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	CheckIfError(err)

	// ... just iterates over the commits
	// ...只是迭代提交
	var cCount int
	err = cIter.ForEach(func(c *object.Commit) error {
		cCount++

		return nil
	})
	CheckIfError(err)

	fmt.Println(cCount)
}
