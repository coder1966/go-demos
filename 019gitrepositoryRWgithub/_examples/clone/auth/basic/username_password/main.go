package main

import (
	"fmt"
	"os"

	git "github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func main() {
	CheckArgs("<url>", "<directory>", "<github_username>", "<github_password>")
	url, directory, username, password := os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	// Clone the given repository to the given directory
	// 将给定存储库克隆到给定目录
	Info("git clone %s %s", url, directory)

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		},
		URL:      url,
		Progress: os.Stdout,
	})
	CheckIfError(err)

	// ... retrieving the branch being pointed by HEAD
	// ...检索 HEAD 指向的分支
	ref, err := r.Head()
	CheckIfError(err)
	// ... retrieving the commit object
	// ...检索提交对象
	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	fmt.Println(commit)
}
