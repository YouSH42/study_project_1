package handler

import (
	"fmt"
	"os"
)

func Error(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
