package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing"
)

// Basic example of how to checkout a specific commit.
// 如何签出特定提交的基本示例。
func main() {
	CheckArgs("<url>", "<directory>", "<commit>")
	url, directory, commit := os.Args[1], os.Args[2], os.Args[3]

	// Clone the given repository to the given directory
	// 将给定存储库克隆到给定目录
	Info("git clone %s %s", url, directory)
	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL: url,
	})

	CheckIfError(err)

	// ... retrieving the commit being pointed by HEAD
	// ...检索 HEAD 指向的提交
	Info("git show-ref --head HEAD")
	ref, err := r.Head()
	CheckIfError(err)
	fmt.Println(ref.Hash())

	w, err := r.Worktree()
	CheckIfError(err)

	// ... checking out to commit
	// ...检查并提交
	Info("git checkout %s", commit)
	err = w.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash(commit),
	})
	CheckIfError(err)

	// ... retrieving the commit being pointed by HEAD, it shows that the repository is pointing to the giving commit in detached mode
	// ...检索 HEAD 指向的提交，它表明存储库指向分离模式下的给定提交
	Info("git show-ref --head HEAD")
	ref, err = r.Head()
	CheckIfError(err)
	fmt.Println(ref.Hash())
}
