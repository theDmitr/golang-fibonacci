package ui

import (
	"fmt"
	"main/service"
	"main/utils"
)

type ConsoleUI struct {
	fibService *service.FibonacciService
}

func NewConsoleUI(fibService *service.FibonacciService) (*ConsoleUI, error) {
	ui := &ConsoleUI{
		fibService: fibService,
	}
	return ui, nil
}

func (ui *ConsoleUI) Start() {
	var msg string = "Enter a int value: "
	input, err := utils.ScanfInteger(&msg)

	if err != nil {
		fmt.Println(err)
		return
	}

	left, right, err := ui.fibService.GetFibonacciNeighbors(input)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Your digit is fibonacci!\nLeft = %d\nRight = %d\n", *left, *right)
}
