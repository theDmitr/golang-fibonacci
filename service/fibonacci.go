package service

import (
	"errors"
	"main/utils"
)

func IsFibonacci(number *int) *bool {
	left := 5**number**number + 4
	right := 5**number**number - 4
	result := *utils.IsPerfectSquare(&left) || *utils.IsPerfectSquare(&right)
	return &result
}

func GetFibonacciNeighbors(number *int) (*int, *int, error) {
	if !*IsFibonacci(number) {
		return nil, nil, errors.New("number isn`t fibonacci")
	}

	prev := *number - 1
	next := *number + 1

	for !*IsFibonacci(&prev) {
		prev--
	}

	for !*IsFibonacci(&next) {
		next++
	}

	return &prev, &next, nil
}
