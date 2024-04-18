package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	HOST = "127.0.0.1"
	PORT = "65456"
)

func main() {
	conn, err := net.Dial("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Println("> connect() failed by error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("> echo-client is activated")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		sendMsg, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, sendMsg)

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("> Error reading from server:", err)
			break
		}
		fmt.Print("> received:", message)

		if sendMsg == "quit\n" {
			break
		}
	}

	fmt.Println("> echo-client is de-activated")
}
