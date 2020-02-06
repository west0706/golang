package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	connect, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	recvBuf := make([]byte, 2048)

	send_msg := "Please Play a game~!"

	for {
		//hearbeat
		connect.Write([]byte(send_msg))
		fmt.Println("Send Data:", send_msg)
		time.Sleep(time.Second * 10)

		n, _ := connect.Read(recvBuf)
		if n > 0 {
			fmt.Println(string(recvBuf))
		}
	}

}
