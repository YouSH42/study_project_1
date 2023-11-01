// 음식 주문 스크립트
var orderForm = document.getElementById("order-form");
// 밥류 관련 스크립트
var foodListRiceButton = document.getElementById("food_list_rice_button");
var foodListRiceCheckbox = document.getElementById("food_list_rice_checkbox");

foodListRiceButton.addEventListener("click", function () {
    if (foodListRiceCheckbox.style.display === "none" || foodListRiceCheckbox.style.display === "") {
        // JSON 데이터를 사용하여 목록 생성
        fetch("/foodricelist")
            .then(response => response.json())
            .then(data => {
                data.forEach(item => {
                    // 리스트 아이템 생성
                    var listItem = document.createElement("li");

                    // 버튼 생성 (주문 버튼)
                    var orderButton = document.createElement("button");
                    orderButton.textContent = "선택";

                    // 라벨 생성 (항목 이름과 가격을 표시)
                    var label = document.createElement("label");
                    label.textContent = `${item.name} - ${item.price}원`;

                    // 아이템을 목록에 추가
                    listItem.appendChild(orderButton);
                    listItem.appendChild(label);
                    foodListRiceCheckbox.appendChild(listItem);

                    // 클릭 이벤트 핸들러를 추가하여 해당 음식 정보를 주문 양식에 채웁니다.
                    orderButton.addEventListener("click", function () {
                        orderForm.elements["name"].value = item.name;
                        orderForm.elements["price"].value = item.price;
                        orderForm.elements["category"].value = item.category;
                    });
                });
            });
        foodListRiceCheckbox.style.display = "block";
    } else {
        foodListRiceCheckbox.style.display = "none";
        //목록 비우기
        while (foodListRiceCheckbox.firstChild) {
            foodListRiceCheckbox.removeChild(foodListRiceCheckbox.firstChild);
        }
    }
});

// 면류 관련 스크립트
var foodListNoodleButton = document.getElementById("food_list_noodle_button");
var foodListNoodleCheckbox = document.getElementById("food_list_noodle_checkbox");

foodListNoodleButton.addEventListener("click", function () {
    if (foodListNoodleCheckbox.style.display === "none" || foodListNoodleCheckbox.style.display === "") {
        // "면류" 데이터를 가져옴
        fetch("/foodnoodlelist")
            .then(response => response.json())
            .then(data => {
                // JSON 데이터를 사용하여 목록 생성
                data.forEach(item => {
                    // 리스트 아이템 생성
                    var listItem = document.createElement("li");

                    // 버튼 생성 (주문 버튼)
                    var orderButton = document.createElement("button");
                    orderButton.textContent = "선택";

                    // 라벨 생성 (항목 이름과 가격을 표시)
                    var label = document.createElement("label");
                    label.textContent = `${item.name} - ${item.price}원`;

                    // 아이템을 목록에 추가
                    listItem.appendChild(orderButton);
                    listItem.appendChild(label);
                    foodListNoodleCheckbox.appendChild(listItem);

                    // 클릭 이벤트 핸들러를 추가하여 해당 음식 정보를 주문 양식에 채웁니다.
                    orderButton.addEventListener("click", function () {
                        orderForm.elements["name"].value = item.name;
                        orderForm.elements["price"].value = item.price;
                        orderForm.elements["category"].value = item.category;
                    });
                });
            });
        foodListNoodleCheckbox.style.display = "block";
    } else {
        foodListNoodleCheckbox.style.display = "none";
        //목록 비우기
        while (foodListNoodleCheckbox.firstChild) {
            foodListNoodleCheckbox.removeChild(foodListNoodleCheckbox.firstChild);
        }
    }
});