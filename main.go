package main

import (
	"main/service"
	"main/ui"
)

func main() {
	var service, _ = service.NewFibonacciService()
	var ui, _ = ui.NewConsoleUI(service)
	ui.Start()
}
