package handler

import "fmt"

func Error(err error) {
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}
