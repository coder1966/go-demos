package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
)

// Example of how to:
// 如何执行的示例：
// - Create a new in-memory repository
// - 创建一个新的内存存储库
// - Create a new remote named "example"
// - 创建一个名为“example”的新远程
// - List remotes and print them
// - 列出遥控器并打印它们
// - Pull using the new remote "example"
// - 使用新的远程“示例”拉取
// - Iterate the references again, but only showing hash references, not symbolic ones
// - 再次迭代引用，但仅显示哈希引用，而不显示符号引用
// - Remove remote "example"
// - 删除远程“示例”
func main() {
	// Create a new repository
	// 创建一个新的存储库
	Info("git init")
	r, err := git.Init(memory.NewStorage(), nil)
	CheckIfError(err)

	// Add a new remote, with the default fetch refspec
	// 添加一个新的遥控器，使用默认的 fetch refspec
	Info("git remote add example https://github.com/git-fixtures/basic.git") // github.com/git-fixtures/basic.git”）
	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: "example",
		URLs: []string{"https://github.com/git-fixtures/basic.git"}, // github.com/git-fixtures/basic.git"},
	})

	CheckIfError(err)

	// List remotes from a repository
	// 列出存储库中的遥控器
	Info("git remote -v")

	list, err := r.Remotes()
	CheckIfError(err)

	for _, r := range list {
		fmt.Println(r)
	}

	// Fetch using the new remote
	// 使用新遥控器获取
	Info("git fetch example")
	err = r.Fetch(&git.FetchOptions{
		RemoteName: "example",
	})

	CheckIfError(err)

	// List the branches > git show-ref
	// 列出分支 > git show-ref
	Info("git show-ref")

	refs, err := r.References()
	CheckIfError(err)

	err = refs.ForEach(func(ref *plumbing.Reference) error {
		// The HEAD is omitted in a `git show-ref` so we ignore the symbolic references, the HEAD
		// HEAD 在 `git show-ref` 中被省略，所以我们忽略符号引用，HEAD
		if ref.Type() == plumbing.SymbolicReference {
			return nil
		}

		fmt.Println(ref)
		return nil
	})

	CheckIfError(err)

	// Delete the example remote
	// 删除示例远程
	Info("git remote rm example")

	err = r.DeleteRemote("example")
	CheckIfError(err)
}
