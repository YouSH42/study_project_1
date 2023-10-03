package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func AddDataHandler(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	Error(err)

	var receivedFoodInfo FoodItem
	err = json.Unmarshal(body, &receivedFoodInfo)
	Error(err)
	fmt.Println("음식 이름 :", receivedFoodInfo.Name, "가격 :", receivedFoodInfo.Price, "카테고리 :", receivedFoodInfo.Category)
	// 데이터베이스 연결
	db, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName)
	Error(err)
	defer db.Close()

	responseMessage := struct {
		Message string `json:"message"`
	}{
		Message: "추가되었습니다",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseJSON, _ := json.Marshal(responseMessage)
	w.Write(responseJSON)

}
