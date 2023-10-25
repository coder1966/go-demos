package main

import (
	"os"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
)

// Example of how to open a repository in a specific path, and push to its default remote (origin).
// 如何在特定路径中打开存储库并将其推送到其默认远程（源）的示例。
func main() {
	CheckArgs("<repository-path>")
	path := os.Args[1]

	r, err := git.PlainOpen(path)
	CheckIfError(err)

	Info("git push")
	// push using default options
	// 使用默认选项推送
	err = r.Push(&git.PushOptions{})
	CheckIfError(err)
}
