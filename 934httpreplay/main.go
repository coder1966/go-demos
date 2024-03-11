package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	protocol = "udp"
	port     = ":8125"
)

func handleConnection(conn *net.UDPConn) {
	var addr *net.UDPAddr
	var n int
	var err error
	buf := make([]byte, 0)       // 最后结果
	buffer := make([]byte, 1024) // 创建一个缓冲区大小为1024的字节数组
	for {
		// 从连接读取数据到缓冲区
		n, addr, err = conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatalf("Error reading from connection: %v", err)
		}

		buf = append(buf, buffer[:n]...)

		// message := string(buf[:]) // 将字节转换成字符串
		// fmt.Println("### Received message: ", message)

	}

	message := string(buf[:]) // 将字节转换成字符串
	fmt.Println("### Received message: ", message)

	response := "Hello, client!"                     // 构造要发送的响应消息
	_, err = conn.WriteToUDP([]byte(response), addr) // 向客户端发送响应消息
	if err != nil {
		log.Fatalf("Error writing to connection: %v", err)
	}
}

func main() {
	// 带任何参数表示回放
	if len(os.Args) > 1 {
		// 回放请求
		// replayRequests()
		fmt.Println("完成了回放----")
		return
	}

	addr, err := net.ResolveUDPAddr(protocol, port) // 设置服务器地址和端口号
	if err != nil {
		log.Fatalf("Failed resolving address: %v", err)
	}
	conn, err := net.ListenUDP("udp", addr) // 开始监听指定地址上的UDP连接
	if err != nil {
		log.Fatalf("Failed listening on address: %v", err)
	}
	defer conn.Close()
	handleConnection(conn) // 处理传入的连接
}
