package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
)

// Graceful cancellation example of a basic git operation such as Clone.
// 基本 git 操作（例如克隆）的优雅取消示例。
func main() {
	CheckArgs("<url>", "<directory>")
	url := os.Args[1]
	directory := os.Args[2]

	// Clone the given repository to the given directory
	// 将给定存储库克隆到给定目录
	Info("git clone %s %s", url, directory)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// The context is the mechanism used by go-git, to support deadlines and cancellation signals.
	// 上下文是 go-git 使用的机制，用于支持截止日期和取消信号。
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers // 当我们完成消耗整数时取消

	go func() {
		<-stop
		Warning("\nSignal detected, canceling operation...")
		cancel()
	}()

	Warning("To gracefully stop the clone operation, push Crtl-C.")

	// Using PlainCloneContext we can provide to a context, if the context is cancelled, the clone operation stops gracefully.
	// 使用 PlainCloneContext 我们可以提供一个上下文，如果上下文被取消，克隆操作就会正常停止。
	_, err := git.PlainCloneContext(ctx, directory, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})

	// If the context was cancelled, an error is returned.
	// 如果上下文被取消，则会返回错误。
	CheckIfError(err)
}
