package main

import (
	"main/service"
	"main/ui"
)

func main() {
	var service *service.FibonacciService = &service.FibonacciService{}
	var ui, _ = ui.NewConsoleUI(service)
	ui.Start()
}
