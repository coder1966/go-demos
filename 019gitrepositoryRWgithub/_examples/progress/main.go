package main

import (
	"os"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
)

// Example of how to show the progress when you do a basic clone operation.
// 执行基本克隆操作时如何显示进度的示例。
func main() {
	CheckArgs("<url>", "<directory>")
	url := os.Args[1]
	directory := os.Args[2]

	// Clone the given repository to the given directory
	// 将给定存储库克隆到给定目录
	Info("git clone %s %s", url, directory)

	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:   url,
		Depth: 1,

		// as git does, when you make a clone, pull or some other operations the server sends information via the sideband, this information can being collected providing a io.Writer to the CloneOptions options
		// 与 git 一样，当您进行克隆、拉取或其他一些操作时，服务器通过边带发送信息，可以通过向 CloneOptions 选项提供 io.Writer 来收集此信息
		Progress: os.Stdout,
	})

	CheckIfError(err)
}
