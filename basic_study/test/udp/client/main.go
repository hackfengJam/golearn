package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "8080", "port")

func Client(name string) {
	flag.Parse()
	addr, err := net.ResolveUDPAddr("udp", *host+":"+*port)
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Can't dial: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	cnt := 1

	go func() {
		for {
			if cnt > 100 {
				break
			}
			_, err = conn.Write([]byte(fmt.Sprintf("jianghaifeng name: %s -> %d", name, cnt)))
			if err != nil {
				fmt.Println("failed:", err)
				os.Exit(1)
			}
			cnt++
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			data := make([]byte, 1024)
			_, err = conn.Read(data)
			if err != nil {
				fmt.Println("failed to read UDP msg because of ", err)
				os.Exit(1)
			}
			fmt.Println(string(data))
			time.Sleep(1 * time.Second)
		}
	}()

	select {}

	// t := binary.BigEndian.Uint32(data)
	// fmt.Println(time.Unix(int64(t), 0).String())
	os.Exit(0)
}

// go run timeclient.go -host time.nist.gov
func main() {
	Client("qingfeng")
}
