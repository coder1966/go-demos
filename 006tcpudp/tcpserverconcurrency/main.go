package main

import (
	"fmt"
	"net"
)

func HandlerConnect(conn net.Conn) {
	// 必须放在子程序，不能放在主程序，否则永远不关闭
	defer conn.Close()

	// server 可以做的一些事情，例如吗，数据处理
	addr := conn.RemoteAddr() // 远端 地址
	fmt.Println("获得socket协议(例如tcp): ", addr.Network())
	fmt.Println("获得客户端ip+port: ", addr.String())

	// 三 循环读取客户端数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)

		if n == 0 {
			fmt.Println(" 客户端stop, 收到字节数 = 0, 断开连接 ")
			return
		}
		msg := string(buf[:n])
		if "exit\n" == msg || "exit\r\n" == msg {
			fmt.Println(" 客户端exit, 断开连接 ")
			return
		}

		if err != nil {
			fmt.Println(" error: ", err) // 注意: 当客户端关闭后,n 等于 0 , err 等于 EOF
			return
		}

		fmt.Println(" 收到", addr.String(), "客户端", n, "字节数据: ", msg)

		// 返回给客户端数据
		_, err = conn.Write([]byte(fmt.Sprintf("收到: %d 字节数据...", n)))
		if err != nil {
			fmt.Println(" error: ", err)
			return
		}

	}

}

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
	for { // 循环等待创建连接，实现可以接收多个client的功能。
		conn, err := listener.Accept() // 等待连接
		if err != nil {
			fmt.Println(" error: ", err)
			return
		}

		// 并发
		go HandlerConnect(conn)
	}

}
