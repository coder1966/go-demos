package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 一 连接UDP服务器
	udpConn, err := net.Dial("udp", "192.168.56.1:8385")
	if err != nil {
		fmt.Println(" error: ", err)
		return
	}
	defer udpConn.Close()
	fmt.Println("UDP客户端启动，请输入消息 ......")

	// 创建携程，获得键盘输入
	go func() {
		str := make([]byte, 4096)
		for {
			n, err := os.Stdin.Read(str) // 获得键盘输入
			if err != nil {
				fmt.Println(" error: ", err)
				continue
			}

			_, err = udpConn.Write(str[:n])
			if err != nil {
				fmt.Println(" error: ", err)
				return
			}
		}
	}()

	// 2 接收 server 端的数据
	buf := make([]byte, 4096)
	for {
		n, err := udpConn.Read(buf) // 返回值 n 是收到的字节数。
		if err != nil {
			fmt.Println(" error: ", err)
			return
		}
		fmt.Println("client 收到数据: ", string(buf[:n])) // 只打印收到的字节数
	}

	//  关闭 ，在 defer 里

}
