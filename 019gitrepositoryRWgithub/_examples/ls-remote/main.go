package main

import (
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/storage/memory"

	. "github.com/go-git/go-git/v5/_examples"
)

// Retrieve remote tags without cloning repository
// 无需克隆存储库即可检索远程标签
func main() {
	CheckArgs("<url>")
	url := os.Args[1]

	Info("git ls-remote --tags %s", url)

	// Create the remote with repository URL
	// 使用存储库 URL 创建远程
	rem := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{url},
	})

	log.Print("Fetching tags...")

	// We can then use every Remote functions to retrieve wanted information
	// 然后我们可以使用每个远程功能来检索想要的信息
	refs, err := rem.List(&git.ListOptions{
		// Returns all references, including peeled references.
		// 返回所有引用，包括剥离的引用。
		PeelingOption: git.AppendPeeled,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Filters the references list and only keeps tags
	// 过滤参考列表并仅保留标签
	var tags []string
	for _, ref := range refs {
		if ref.Name().IsTag() {
			tags = append(tags, ref.Name().Short())
		}
	}

	if len(tags) == 0 {
		log.Println("No tags!")
		return
	}

	log.Printf("Tags found: %v", tags)
}
