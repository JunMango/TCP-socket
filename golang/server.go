package main

import (
	"fmt"
	"net"
	"os"
)

const (
	HOST = "127.0.0.1"
	PORT = "65456"
)

func main() {
	// TCP 서버 소켓 생성
	listener, err := net.Listen("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Println("> Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("> echo-server is activated")

	// 클라이언트의 연결을 대기하고, 연결이 되면 새로운 고루틴으로 처리
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("> Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 클라이언트의 정보 출력
	fmt.Printf("> client connected by IP address %s with Port number %d\n", conn.RemoteAddr().(*net.TCPAddr).IP, conn.RemoteAddr().(*net.TCPAddr).Port)

	// 클라이언트와의 통신 루프
	for {
		// 데이터 수신
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println("> Error reading:", err)
			return
		}
		// 수신한 데이터 출력
		fmt.Println("> echoed:", string(data[:n]))
		// 클라이언트에게 데이터 전송
		_, err = conn.Write(data[:n])
		if err != nil {
			fmt.Println("> Error writing:", err)
			return
		}
		// 클라이언트가 'quit'을 보내면 연결 종료
		if string(data[:n]) == "quit" {
			break
		}
	}
}
