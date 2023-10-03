package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	Error(err)

	var receivedFoodInfo FoodItem
	err = json.Unmarshal(body, &receivedFoodInfo)
	Error(err)
	fmt.Println("음식 이름 : ", receivedFoodInfo.FoodName, "가격 : ", receivedFoodInfo.Price, "카테고리", receivedFoodInfo.Category)
}
