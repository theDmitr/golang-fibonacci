package ui

import (
	"fmt"
	"main/service"
	"main/utils"
)

func Start() {
	var msg string = "Enter a int value: "
	input, err := utils.ScanfInteger(&msg)

	if err != nil {
		fmt.Println(err)
		return
	}

	left, right, err := service.GetFibonacciNeighbors(input)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Your digit is fibonacci!\nLeft = %d\nRight = %d\n", *left, *right)
}
