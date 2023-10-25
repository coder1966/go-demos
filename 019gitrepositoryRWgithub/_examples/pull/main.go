package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
)

// Pull changes from a remote repository
// 从远程存储库中提取更改
func main() {
	CheckArgs("<path>")
	path := os.Args[1]

	// We instantiate a new repository targeting the given path (the .git folder)
	// 我们实例化一个针对给定路径（.git 文件夹）的新存储库
	r, err := git.PlainOpen(path)
	CheckIfError(err)

	// Get the working directory for the repository
	// 获取存储库的工作目录
	w, err := r.Worktree()
	CheckIfError(err)

	// Pull the latest changes from the origin remote and merge into the current branch
	// 从原始远程拉取最新更改并合并到当前分支
	Info("git pull origin")
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	CheckIfError(err)

	// Print the latest commit that was just pulled
	// 打印刚刚拉取的最新提交
	ref, err := r.Head()
	CheckIfError(err)
	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	fmt.Println(commit)
}
