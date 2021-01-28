package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "8080", "port")

// go run timeclient.go -host time.nist.gov
func main() {
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
	for {
		if cnt > 100 {
			break
		}
		_, err = conn.Write([]byte(fmt.Sprintf("jianghaifeng -> %d", cnt)))
		if err != nil {
			fmt.Println("failed:", err)
			os.Exit(1)
		}
		cnt++
		// time.Sleep(1 * time.Second)
	}

	data := make([]byte, 1024)
	_, err = conn.Read(data)
	if err != nil {
		fmt.Println("failed to read UDP msg because of ", err)
		os.Exit(1)
	}
	// t := binary.BigEndian.Uint32(data)
	// fmt.Println(time.Unix(int64(t), 0).String())
	fmt.Println(string(data))
	os.Exit(0)
}
