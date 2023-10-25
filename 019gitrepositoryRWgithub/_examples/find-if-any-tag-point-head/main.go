package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing"
)

// Basic example of how to find if HEAD is tagged.
// 如何查找 HEAD 是否已标记的基本示例。
func main() {
	CheckArgs("<path>")
	path := os.Args[1]

	// We instantiate a new repository targeting the given path (the .git folder)
	// 我们实例化一个针对给定路径（.git 文件夹）的新存储库
	r, err := git.PlainOpen(path)
	CheckIfError(err)

	// Get HEAD reference to use for comparison later on.
	// 获取 HEAD 参考以供稍后进行比较。
	ref, err := r.Head()
	CheckIfError(err)

	tags, err := r.Tags()
	CheckIfError(err)

	// List all tags, both lightweight tags and annotated tags and see if some tag points to HEAD reference.
	// 列出所有标签，包括轻量级标签和带注释的标签，并查看某些标签是否指向 HEAD 引用。
	err = tags.ForEach(func(t *plumbing.Reference) error {
		// This technique should work for both lightweight and annotated tags.
		// 该技术应该适用于轻量级标签和带注释的标签。
		revHash, err := r.ResolveRevision(plumbing.Revision(t.Name()))
		CheckIfError(err)
		if *revHash == ref.Hash() {
			fmt.Printf("Found tag %s with hash %s pointing to HEAD %s\n", t.Name().Short(), revHash, ref.Hash())
		}
		return nil
	})
}
