package service

import (
	"errors"
	"main/utils"
)

type FibonacciService struct {
}

func (service *FibonacciService) IsFibonacci(number *int) *bool {
	left := 5**number**number + 4
	right := 5**number**number - 4
	result := *utils.IsPerfectSquare(&left) || *utils.IsPerfectSquare(&right)
	return &result
}

func (service *FibonacciService) GetFibonacciNeighbors(number *int) (*int, *int, error) {
	if !*service.IsFibonacci(number) {
		return nil, nil, errors.New("number isn`t fibonacci")
	}

	prev := *number - 1
	next := *number + 1

	for !*service.IsFibonacci(&prev) {
		prev--
	}

	for !*service.IsFibonacci(&next) {
		next++
	}

	return &prev, &next, nil
}
