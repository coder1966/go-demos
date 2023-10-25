package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/transport/client"
	githttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
)

// Here is an example to configure http client according to our own needs.
// 下面是一个根据我们自己的需求配置http客户端的例子。
func main() {
	CheckArgs("<url>")
	url := os.Args[1]

	// Create a custom http(s) client with your config
	// 使用您的配置创建自定义 http(s) 客户端
	customClient := &http.Client{
		// accept any certificate (might be useful for testing)
		// 接受任何证书（可能对测试有用）
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},

		// 15 second timeout
		// 15秒超时
		Timeout: 15 * time.Second,

		// don't follow redirect
		// 不遵循重定向
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// Override http(s) default protocol to use our custom client
	// 覆盖 http(s) 默认协议以使用我们的自定义客户端
	client.InstallProtocol("https", githttp.NewClient(customClient))

	// Clone repository using the new client if the protocol is https://
	// 如果协议是 https://，则使用新客户端克隆存储库
	Info("git clone %s", url)

	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: url})
	CheckIfError(err)

	// Retrieve the branch pointed by HEAD
	// 检索 HEAD 指向的分支
	Info("git rev-parse HEAD")

	head, err := r.Head()
	CheckIfError(err)
	fmt.Println(head.Hash())
}
