package main

import (
	"fmt"
	"os"

	git "github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/protocol/packp/capability"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func main() {
	CheckArgs("<url>", "<directory>", "<azuredevops_username>", "<azuredevops_password>")
	url, directory, username, password := os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	// Clone the given repository to the given directory
	// 将给定存储库克隆到给定目录
	Info("git clone %s %s", url, directory)

	// Azure DevOps requires capabilities multi_ack / multi_ack_detailed, which are not fully implemented and by default are included in transport.UnsupportedCapabilities.
	// Azure DevOps 需要 multi_ack / multi_ack_detailed 功能，这些功能尚未完全实现，默认情况下包含在 Transport.UnsupportedCapability 中。
	//
	// The initial clone operations require a full download of the repository, and therefore those unsupported capabilities are not as crucial, so by removing them from that list allows for the first clone to work successfully.
	// 初始克隆操作需要完整下载存储库，因此那些不受支持的功能并不那么重要，因此通过从该列表中删除它们可以使第一个克隆成功工作。
	//
	// Additional fetches will yield issues, therefore work always from a clean clone until those capabilities are fully supported.
	// 额外的提取会产生问题，因此始终从干净的克隆开始工作，直到完全支持这些功能。
	//
	// New commits and pushes against a remote worked without any issues.
	// 针对远程的新提交和推送工作没有任何问题。
	transport.UnsupportedCapabilities = []capability.Capability{
		capability.ThinPack,
	}

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
