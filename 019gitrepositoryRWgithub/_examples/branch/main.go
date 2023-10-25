package main

import (
	"os"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing"
)

// An example of how to create and remove branches or any other kind of reference.
// 如何创建和删除分支或任何其他类型的引用的示例。
func main() {
	CheckArgs("<url>", "<directory>")
	url, directory := os.Args[1], os.Args[2]

	// Clone the given repository to the given directory
	// 将给定存储库克隆到给定目录
	Info("git clone %s %s", url, directory)
	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL: url,
	})
	CheckIfError(err)

	// Create a new branch to the current HEAD
	// 在当前 HEAD 上创建一个新分支
	Info("git branch my-branch")

	headRef, err := r.Head()
	CheckIfError(err)

	// Create a new plumbing.HashReference object with the name of the branch and the hash from the HEAD. The reference name should be a full reference name and not an abbreviated one, as is used on the git cli.
	// 使用分支名称和 HEAD 中的哈希值创建一个新的 plumbing.HashReference 对象。引用名称应该是完整的引用名称，而不是 git cli 上使用的缩写名称。
	//
	// For tags we should use `refs/tags/%s` instead of `refs/heads/%s` used for branches.
	// 对于标签，我们应该使用 `refs/tags/%s` 而不是用于分支的 `refs/heads/%s`。
	ref := plumbing.NewHashReference("refs/heads/my-branch", headRef.Hash())

	// The created reference is saved in the storage.
	// 创建的参考保存在存储器中。
	err = r.Storer.SetReference(ref)
	CheckIfError(err)

	// Or deleted from it.
	// 或者从中删除。
	Info("git branch -D my-branch")
	err = r.Storer.RemoveReference(ref.Name())
	CheckIfError(err)
}
