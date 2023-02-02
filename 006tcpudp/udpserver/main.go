package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 一 组织一个udp地址，绑定 IP+port
	serverAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:8385")
	if err != nil {
		fmt.Println(" error: ", err) // 当server断开与client的连接后 client read err : 0 EOF
		return
	}

	// 二 创建用于通信的 socket
	// 注意udp只有一个套接字socket
	udpConn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		fmt.Println(" error: ", err) // 当server断开与client的连接后 client read err : 0 EOF
		return
	}
	defer udpConn.Close()
	fmt.Println(" UDP server 启动 ")

	// 三 读取 client 发的数据
	buf := make([]byte, 4096)
	n, clientAddr, err := udpConn.ReadFromUDP(buf) // 返回值 n 是收到的字节数。clientAddr客户端地址
	if err != nil {
		fmt.Println(" error: ", err) // 当server断开与client的连接后 client read err : 0 EOF
		return
	}

	// 四 处理数据
	fmt.Println(" 收到", clientAddr, "客户端", n, "字节数据: ", string(buf[:n]))

	// 回应客户端
	nowTime := time.Now().String()
	n, err = udpConn.WriteToUDP([]byte("收到："+nowTime), clientAddr)
	if err != nil {
		fmt.Println(" error: ", err)
		return
	}
	_ = n

}
