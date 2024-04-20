package utils

import (
	"fmt"
	"math"
)

func IsPerfectSquare(x *int) *bool {
	sqrt := int(math.Sqrt(float64(*x)))
	result := sqrt*sqrt == *x
	return &result
}

func ScanfInteger(msg *string) (*int, error) {
	var d int

	fmt.Print(*msg)
	_, err := fmt.Scanf("%d", &d)

	if err != nil {
		return nil, err
	}

	return &d, err
}
