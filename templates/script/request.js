
// 밥류 관련 스크립트
var foodListRiceButton = document.getElementById("food_list_rice_button");
var foodListRice = document.getElementById("food_list_rice");

foodListRiceButton.addEventListener("click", function () {
    if (foodListRice.style.display === "none" || foodListRice.style.display === "") {
        fetch("/foodricelist") // 서버의 데이터 엔드포인트로 요청 보내기
            .then(response => response.json())
            .then(data => {
                // JSON 데이터를 사용하여 목록 생성
                data.forEach(item => {
                    var listItem = document.createElement("li");
                    listItem.textContent = `${item.name} - ${item.price}원 (카테고리: ${item.category})`;
                    foodListRice.appendChild(listItem);
                });
            });
        foodListRice.style.display = "block";
    } else {
        foodListRice.style.display = "none";
        // 목록 비우기
        while (foodListRice.firstChild) {
            foodListRice.removeChild(foodListRice.firstChild);
        }
    }
});
// 면류 관련 스크립트
var foodListNoodleButton = document.getElementById("food_list_noodle_button");
var foodListNoodle = document.getElementById("food_list_noodle");

foodListNoodleButton.addEventListener("click", function () {
    if (foodListNoodle.style.display === "none" || foodListNoodle.style.display === "") {
        fetch("/foodnoodlelist") // 서버의 데이터 엔드포인트로 요청 보내기
            .then(response => response.json())
            .then(data => {
                // JSON 데이터를 사용하여 목록 생성
                data.forEach(item => {
                    var listItem = document.createElement("li");
                    listItem.textContent = `${item.name} - ${item.price}원 (카테고리: ${item.category})`;
                    foodListNoodle.appendChild(listItem);
                });
            });
        foodListNoodle.style.display = "block";
    } else {
        foodListNoodle.style.display = "none";
        // 목록 비우기
        while (foodListNoodle.firstChild) {
            foodListNoodle.removeChild(foodListNoodle.firstChild);
        }
    }
});
//음식 등록 스크립트
document.addEventListener("DOMContentLoaded", function () {
    const form = document.getElementById("food_form");
    const responseDiv = document.getElementById("response");

    form.addEventListener("submit", function (event) {
        event.preventDefault();

        //사용자로부터 입력받은 데이터를 가져와 JSON 형식으로 생성
        const name = document.getElementById("name").value;
        const price = parseInt(document.getElementById("price").value);
        const category = document.getElementById("category").value;
        const jsonData = JSON.stringify({ name, price, category });
        console.log(name);
        console.log(price);
        console.log(category);

        //서버로 JSON 데이터 전송
        fetch("/registration", {
            method: "POST",
            body: jsonData,
            headers: {
                "Content-Type": "application/json",
            }
        })
            .then((response) => response.json())
            .then((data) => {
                // 서버에서 받은 응답을 출력
                responseDiv.textContent = `서버에서 받은 응답: ${JSON.stringify(data)}`;
            })
            .catch((error) => {
                console.error("에러:", error);
            });
    });
});

