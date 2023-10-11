package handler

import (
	"fmt"
	"net"

	_ "github.com/go-sql-driver/mysql"
)

func OrderListener() {
	// TCP 서버 소켓을 열고 포트에서 들어오는 연결을 수신합니다.
	listener, err := net.Listen("tcp", ":9000")
	Error(err)
	defer listener.Close()
	fmt.Println("task가 9000 포트에서 실행 중입니다.(tcp/ip 통신)")

	for {
		// 클라이언트 연결을 대기하고 연결이 들어오면 핸들러 함수를 호출합니다.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("에러:", err)
			continue
		}
		go WriteOrderHandler(conn)
	}
}
