package main

import (
	"fmt"
	"net"
)

// UDP 客户端
func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		//IP:   net.IPv4(10, 250, 212, 255),
		IP:   net.IPv4(255, 255, 255, 255),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接UDP服务器失败，err: ", err)
		return
	}
	defer socket.Close()
	sendData := []byte("Hello Server")
	_, err = socket.Write(sendData) // 发送数据
	if err != nil {
		fmt.Println("发送数据失败，err: ", err)
		return
	}
}
