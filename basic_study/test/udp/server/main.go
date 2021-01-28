package main

import (
	"fmt"
	"net"
	"os"
)

type Body struct {
	Data    []byte
	UDPAddr *net.UDPAddr
}

// 限制goroutine数量
var limitChan = make(chan bool, 1000)
var limitSendChan = make(chan bool, 1000)
var d = make(chan Body, 1000)

// UDP goroutine 实现并发读取UDP数据
func udpProcess(conn *net.UDPConn) {

	// 最大读取数据大小
	data := make([]byte, 1024)
	n, uDPAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("failed read udp msg, error: " + err.Error())
	}
	str := string(data[:n])
	fmt.Println("receive from client, data:" + str)
	d <- Body{Data: data, UDPAddr: uDPAddr}
	<-limitSendChan
}

// UDP goroutine 实现并发读取UDP数据
func udpProcessSend(conn *net.UDPConn) {

	data := <-d
	// 最大读取数据大小
	_, err := conn.WriteToUDP([]byte(fmt.Sprintf("Server -> %s", string(data.Data))), data.UDPAddr)
	if err != nil {
		fmt.Println("failed read udp msg, error: " + err.Error())
	}

	<-limitChan
}

func udpServer(address string) {
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	conn, err := net.ListenUDP("udp", udpAddr)
	defer conn.Close()
	if err != nil {
		fmt.Println("read from connect failed, err:" + err.Error())
		os.Exit(1)
	}

	go func() {
		for {
			limitChan <- true
			go udpProcess(conn)
		}
	}()
	go func() {
		for {
			limitSendChan <- true
			go udpProcessSend(conn)
		}
	}()
	select {}
}

func main() {
	address := "0.0.0.0:8080"
	udpServer(address)
}
