package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/format/config"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// This example requires building with the sha256 tag for it to work: go run -tags sha256 main.go /tmp/repository
// 此示例需要使用 sha256 标签进行构建才能正常工作： go run -tags sha256 main.go /tmp/repository

// Basic example of how to initialise a repository using sha256 as the hashing algorithmn.
// 如何使用 sha256 作为哈希算法初始化存储库的基本示例。
func main() {
	CheckArgs("<directory>")
	directory := os.Args[1]

	os.RemoveAll(directory)

	// Init a new repository using the ObjectFormat SHA256.
	// 使用 ObjectFormat SHA256 初始化新存储库。
	r, err := git.PlainInitWithOptions(directory, &git.PlainInitOptions{ObjectFormat: config.SHA256})
	CheckIfError(err)

	w, err := r.Worktree()
	CheckIfError(err)

	// ... we need a file to commit so let's create a new file inside of the worktree of the project using the go standard library.
	// ...我们需要一个文件来提交，所以让我们使用 go 标准库在项目的工作树内创建一个新文件。
	Info("echo \"hello world!\" > example-git-file")
	filename := filepath.Join(directory, "example-git-file")
	err = os.WriteFile(filename, []byte("hello world!"), 0644)
	CheckIfError(err)

	// Adds the new file to the staging area.
	// 将新文件添加到暂存区域。
	Info("git add example-git-file")
	_, err = w.Add("example-git-file")
	CheckIfError(err)

	// Commits the current staging area to the repository, with the new file just created. We should provide the object.Signature of Author of the commit Since version 5.0.1, we can omit the Author signature, being read from the git config files.
	// 将当前暂存区域以及刚刚创建的新文件提交到存储库。我们应该提供对象。提交的作者签名从版本 5.0.1 开始，我们可以省略从 git 配置文件中读取的作者签名。
	Info("git commit -m \"example go-git commit\"")
	commit, err := w.Commit("example go-git commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "John Doe",
			Email: "john@doe.org",
			When:  time.Now(),
		},
	})

	CheckIfError(err)

	// Prints the current HEAD to verify that all worked well.
	// 打印当前的 HEAD 以验证一切正常。
	Info("git show -s")
	obj, err := r.CommitObject(commit)
	CheckIfError(err)

	fmt.Println(obj)
}