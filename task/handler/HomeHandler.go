package handler

import (
	"html/template"
	"net/http"
)

// 데이터 베이스 연결 정보
const (
	dbDriver   = "mysql"
	dbUser     = "root"
	dbPassword = "4613"
	dbName     = "task"
)

// 음식 정보 저장
type FoodItem struct {
	FoodName string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	// HTML 템플릿 렌더링
	tmpl := template.Must(template.ParseFiles("templates/task.html"))
	tmpl.Execute(w, nil)
}
