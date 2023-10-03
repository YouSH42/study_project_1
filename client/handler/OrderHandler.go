package handler

import (
	"encoding/json"
	"net"
	"net/http"
)

func OrderHandler(w http.ResponseWriter, r *http.Request) {

	// TCP 서버에 연결
	conn, err := net.Dial("tcp", "localhost:9000")
	Error(err)
	defer conn.Close()
	// JSON 데이터 생성
	fooditem := FoodItem{FoodName: "Guksoo", Price: 2500, Category: "C2"}
	jsonData, err := json.Marshal(fooditem)
	Error(err)
	// JSON 데이터를 서버로 전송
	_, err = conn.Write(jsonData)
	Error(err)
}
