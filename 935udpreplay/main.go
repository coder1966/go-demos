package main

// 编译 go build
// 录制 ./main r  （实际上，跟任何字符串都行）  录好了，ctrl+c 退出
// 回放 ./main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	protocol = "udp"
	port     = ":8125"
	fileName = "outFile"
)

var byteCh = make(chan []byte, 1024)

func handleSignal(ch chan os.Signal) {
	// 等待接收到中断信号
	sig := <-ch
	switch sig {
	case syscall.SIGINT: // Ctrl+C
		fallthrough
	case syscall.SIGTERM: // kill命令或者kill -15
		fmt.Println("Received interrupt signal")
		// 这里可以编写清理资源、关闭连接等操作
	}

	close(byteCh)
	data := []byte{}
	for v := range byteCh {
		data = append(data, v...)
	}
	err := os.WriteFile(fileName, data, 0666) //写入文件(字节数组)
	if err != nil {
		panic(err)
	}

	os.Exit(0)
}

func handleConnection(conn *net.UDPConn) {

	for {
		buffer := make([]byte, 1024*1024) // 创建一个缓冲区大小为1024的字节数组
		// 从连接读取数据到缓冲区
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error reading from connection: %v", err)
		}

		byteCh <- buffer[:n]

		message := string(buffer[:n]) // 将字节转换成字符串
		fmt.Println("### Received message: ", message)

		response := "Hello, client!"                     // 构造要发送的响应消息
		_, err = conn.WriteToUDP([]byte(response), addr) // 向客户端发送响应消息
		if err != nil {
			fmt.Printf("Error writing to connection: %v", err)
		}
	}
}

func replayRequests() {
	conn, err := net.Dial(protocol, "127.0.0.1"+port)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 要发送的消息内容
	message, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	_, err = conn.Write(message)
	if err != nil {
		panic(err)
	}

	fmt.Println("### replay message: ", string(message))

}

func main() {
	// 不带任何参数表示回放
	if len(os.Args) < 2 {
		// 回放请求
		replayRequests()
		fmt.Println("完成了回放----")
		return
	}

	// 设置信号处理程序
	c := make(chan os.Signal, 2)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go handleSignal(c)

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
