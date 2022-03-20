package main

//	udp 聊天室
import (
	"fmt"
	"github.com/axgle/mahonia"
	"net"
	"strconv"
	"strings"
	"sync"
)

var (
	name string
	a    byte
	b    byte
	c    byte
	d    byte
)

func lt(wg *sync.WaitGroup) {

	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(a, b, c, d),
		Port: 30000,
	})
	defer listen.Close()
	if err != nil {
		fmt.Println("监听端口错误: ", err)
		return
	}
	enc := mahonia.NewEncoder("utf8")
	for {
		var data [204800]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("读取消息错误 ", err)
			continue
		}
		fmt.Println(addr, enc.ConvertString(string(data[:n])), "\r\n")
	}
	wg.Done()
}
func fs(wg *sync.WaitGroup) {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		//IP:   net.IPv4(10, 250, 212, 255),
		IP:   net.IPv4(a, b, c, 255),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接UDP服务器失败，err: ", err)
		return
	}
	defer socket.Close()
	var s string
	for {
		fmt.Scan(&s)
		s := name + ":\r\n" + s
		sendData := []byte(s)
		_, err = socket.Write(sendData) // 发送数据
		if err != nil {
			fmt.Println("发送数据失败，err: ", err)
			continue
		}
	}
	wg.Done()
}

func select_list() int {
	var list int
	fmt.Println(`
	请输入功能序号：
	1.进入群聊
`)
	fmt.Scan(&list)
	return list
}
func main() {
	conn, err := net.Dial("udp", "www.google.com.hk:80")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	ip := strings.Split(conn.LocalAddr().String(), ":")[0]
	ipp := strings.Split(ip, ".")
	aa, err := strconv.Atoi(ipp[0])
	bb, err := strconv.Atoi(ipp[1])
	cc, err := strconv.Atoi(ipp[2])
	dd, err := strconv.Atoi(ipp[3])
	a = byte(aa)
	b = byte(bb)
	c = byte(cc)
	d = byte(dd)

	var wg sync.WaitGroup
	fmt.Println("大佬，给自己起个别名吧，请输入：")
	fmt.Scan(&name)
	fmt.Println("您已获得屌炸天的昵称：", name)
	list := select_list()
	switch list {
	case 1:
		fmt.Println("已进入群聊，起飞！")
		go lt(&wg)
		go fs(&wg)
		wg.Add(2)
	}
	wg.Wait()

}
