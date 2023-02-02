package main

import (
	"fmt"
	"net"
)

func main() {
	// 一 创建套接字 socket
	// 指定 server 端使用的通信协议，绑定 IP+port
	listener, err := net.Listen("tcp", "0.0.0.0:8384")
	if err != nil {
		fmt.Println(" error: ", err)
		return
	}
	defer listener.Close()
	fmt.Println("server start success, waiting connect ......")

	// 二 创建用于连接和通信的 socket
	// 调用 Accept 监听连接，此时会注射 server 端，直到有客户发送连接。
	// server 端每收到一个连接，都会传建一个conn 套接字(socket)。
	conn, err := listener.Accept() // 等待连接
	if err != nil {
		fmt.Println(" error: ", err)
		return
	}
	defer conn.Close()

	// 三 向 client 收发数据
	// 1 读取客户端发送的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf) // 返回值 n 是收到的字节数。
	if err != nil {
		fmt.Println(" error: ", err)
		return
	}
	fmt.Println("收到数据: ", string(buf[:n])) // 只打印收到的字节数

	// 2 向 client 发送数据
	_, err = conn.Write([]byte(fmt.Sprintf("收到: %d 字节数据...", n)))
	if err != nil {
		fmt.Println(" error: ", err)
		return
	}

	// 四 关闭 ，在 2 个 defer 里

}
