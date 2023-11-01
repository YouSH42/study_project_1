package handler

import (
	"encoding/json"
	"net"
	"net/http"
	"strconv"
)

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	//URL 쿼리 매개변수 가져오기
	queryParams := r.URL.Query()

	name := queryParams.Get("name")
	price := queryParams.Get("price")
	category := queryParams.Get("category")
	intPrice, err := strconv.Atoi(price)
	Error(err)
	fooditem := FoodItem{Name: name, Price: intPrice, Category: category}

	// TCP 서버에 연결
	conn, err := net.Dial("tcp", "localhost:9000")
	Error(err)
	defer conn.Close()
	// JSON 데이터 생성
	jsonData, err := json.Marshal(fooditem)
	Error(err)
	// JSON 데이터를 서버로 전송
	_, err = conn.Write(jsonData)
	Error(err)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
