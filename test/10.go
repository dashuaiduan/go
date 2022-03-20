package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	defer listen.Close()
	if err != nil {
		fmt.Println("Listen failed, err: ", err)
		return
	}
	for {
		var data [2048]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("read udp failed, err: ", err)
			continue
		}
		fmt.Println("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		//_, err = listen.WriteToUDP(data[:n], addr) // 发送数据
		//if err != nil {
		//	fmt.Println("Write to udp failed, err: ", err)
		//	continue
		//}
	}

}
