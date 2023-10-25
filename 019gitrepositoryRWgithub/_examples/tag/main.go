package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Basic example of how to list tags.
// 如何列出标签的基本示例。
func main() {
	CheckArgs("<path>")
	path := os.Args[1]

	// We instantiate a new repository targeting the given path (the .git folder)
	// 我们实例化一个针对给定路径（.git 文件夹）的新存储库
	r, err := git.PlainOpen(path)
	CheckIfError(err)

	// List all tag references, both lightweight tags and annotated tags
	// 列出所有标签引用，包括轻量级标签和带注释的标签
	Info("git show-ref --tag")

	tagrefs, err := r.Tags()
	CheckIfError(err)
	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		fmt.Println(t)
		return nil
	})
	CheckIfError(err)

	// Print each annotated tag object (lightweight tags are not included)
	// 打印每个带注释的标签对象（不包括轻量级标签）
	Info("for t in $(git show-ref --tag); do if [ \"$(git cat-file -t $t)\" = \"tag\" ]; then git cat-file -p $t ; fi; done")

	tags, err := r.TagObjects()
	CheckIfError(err)
	err = tags.ForEach(func(t *object.Tag) error {
		fmt.Println(t)
		return nil
	})
	CheckIfError(err)
}
