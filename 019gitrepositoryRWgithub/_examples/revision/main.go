package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing"
)

// Example how to resolve a revision into its commit counterpart
// 如何将修订解析为其对应的提交的示例
func main() {
	CheckArgs("<path>", "<revision>")

	path := os.Args[1]
	revision := os.Args[2]

	// We instantiate a new repository targeting the given path (the .git folder)
	// 我们实例化一个针对给定路径（.git 文件夹）的新存储库
	r, err := git.PlainOpen(path)
	CheckIfError(err)

	// Resolve revision into a sha1 commit, only some revisions are resolved look at the doc to get more details
	// 将修订解决为 sha1 提交，仅解决了部分修订，请查看文档以获取更多详细信息
	Info("git rev-parse %s", revision)

	h, err := r.ResolveRevision(plumbing.Revision(revision))

	CheckIfError(err)

	fmt.Println(h.String())
}
