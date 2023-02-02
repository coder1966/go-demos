package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 一 客户端创建用于通信的 socket
	// 指定通信协议，指定server 的 IP+port
	conn, err := net.Dial("tcp", "192.168.56.1:8384")
	if err != nil {
		fmt.Println(" error: ", err)
		return
	}
	defer conn.Close()
	fmt.Println("客户端启动，请输入消息 ......")

	// 创建携程，获得键盘舒缓乳
	go func() {
		str := make([]byte, 4096)
		for {
			n, err := os.Stdin.Read(str) // 获得键盘输入
			if err != nil {
				fmt.Println(" error: ", err)
				continue
			}

			_, err = conn.Write(str[:n])
			if err != nil {
				fmt.Println(" error: ", err)
				return
			}
		}
	}()

	// 2 接收 server 端的数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf) // 返回值 n 是收到的字节数。
		if err != nil {
			fmt.Println(" error: ", err)
			return
		}
		fmt.Println("client 收到数据: ", string(buf[:n])) // 只打印收到的字节数
	}

	//  关闭 ，在 defer 里

}

/*
func main() {
	// 一 客户端创建用于通信的 socket
	// 指定通信协议，指定server 的 IP+port
	conn, err := net.Dial("tcp", "192.168.56.1:8384")
	if err != nil {
		fmt.Println(" error: ", err)
		return
	}
	defer conn.Close()
	fmt.Println("客户端启动 ......")

	// 二 收发 server 数据
	// 1 向 server 发数据
	msg := "你好吗？ARE YOU OK?"
	_, err = conn.Write([]byte(msg))

	if err != nil {
		fmt.Println(" error: ", err)
		return
	}

	// 2 接收 server 端的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf) // 返回值 n 是收到的字节数。
	if err != nil {
		fmt.Println(" error: ", err)
		return
	}
	fmt.Println("client 收到数据: ", string(buf[:n])) // 只打印收到的字节数

	//  关闭 ，在 defer 里

} */
