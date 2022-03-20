package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("udp", "www.google.com.hk:80")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	ip := strings.Split(conn.LocalAddr().String(), ":")[0]
	fmt.Println(ip)
	a := strings.Split(ip, ".")
	fmt.Println(a)

	fmt.Println(byte(11))
}
