package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// connect, err := net.Dial("tcp", "127.0.0.1:5555")
	connect, err := net.Dial("tcp", "172.17.0.4:5555")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		//hearbeat
		connect.Write([]byte("HIHI"))
		fmt.Println("Send Data:", "HIHI")
		time.Sleep(time.Second * 1)
	}

}
