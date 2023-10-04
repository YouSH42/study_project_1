package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type FoodItem struct {
	FoodName string `json:"foodName"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

func main() {

	// TCP 서버 소켓을 열고 포트에서 들어오는 연결을 수신합니다.
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("에러:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("서버가 9000 포트에서 실행 중입니다.")

	for {
		// 클라이언트 연결을 대기하고 연결이 들어오면 핸들러 함수를 호출합니다.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("에러:", err)
			continue
		}
		go handleClient(conn)
	}
	//router := mux.NewRouter()
	//router.HandleFunc("/", handler.Home).Methods("GET")

	//log.Println("Listening on :9000")
	//log.Fatal(http.ListenAndServe(":9000", router))
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	fmt.Println("클라이언트 연결 수락됨")

	// JSON 데이터 수신
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("에러:", err)
		return
	}

	// 수신된 JSON 데이터 파싱
	var receivedFoodItem FoodItem
	err = json.Unmarshal(buffer[:n], &receivedFoodItem)
	if err != nil {
		fmt.Println("에러:", err)
		return
	}

	// 파싱된 데이터 출력
	fmt.Printf("클라이언트에서 받은 JSON 데이터:\n%+v\n", receivedFoodItem)
}
