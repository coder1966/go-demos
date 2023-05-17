package main

/*

 */

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	osSignal := make(chan os.Signal, 1)
	exitCh := make(chan struct{}, 1)
	// 监控哪些信号。（就是 kill -? 进程号；例如 kill -1 XXX）
	signal.Notify(osSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	go waitOsSignal(osSignal, exitCh)
	fmt.Println("阻塞在这里，等待60秒结束")
	time.Sleep(time.Second * 60)

	fmt.Println("main 结束")
}

func waitOsSignal(osSignal chan os.Signal, exitCh chan struct{}) {
	fmt.Println("等待系统信号")
	// 阻塞在这里
	s := <-osSignal
	// 给全体广播信号
	close(exitCh)
	fmt.Println("收到系统信号： ", s)
}
