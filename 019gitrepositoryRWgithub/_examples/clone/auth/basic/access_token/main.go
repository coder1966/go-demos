package main

import (
	"fmt"
	"os"

	git "github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func main() {
	CheckArgs("<url>", "<directory>", "<github_access_token>")
	url, directory, token := os.Args[1], os.Args[2], os.Args[3]

	// Clone the given repository to the given directory
	// 将给定存储库克隆到给定目录
	Info("git clone %s %s", url, directory)

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		// The intended use of a GitHub personal access token is in replace of your password because access tokens can easily be revoked.
		// GitHub 个人访问令牌的预期用途是代替您的密码，因为访问令牌很容易被撤销。
		// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
		Auth: &http.BasicAuth{
			Username: "abc123", // yes, this can be anything except an empty string // 是的，这可以是除空字符串之外的任何内容
			Password: token,
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
