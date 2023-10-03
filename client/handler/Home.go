package handler

import (
	"html/template"
	"net/http"
)

// 데이터베이스 연결 정보
const (
	dbDriver   = "mysql"
	dbUser     = "root"
	dbPassword = "enddl4613"
	dbName     = "kiosk"
)

// 데이터를 저장할 구조체 정의
type FoodItem struct {
	FoodName string `json:"foodName"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

func Home(w http.ResponseWriter, _ *http.Request) {
	// HTML 템플릿 렌더링
	tmpl := template.Must(template.ParseFiles("templates/client.html"))
	tmpl.Execute(w, nil)
}
