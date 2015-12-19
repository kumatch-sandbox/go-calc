package calc

import (
	"gopkg.in/kumatch-sandbox/go-arithmetic.v1"
)

func Add(a, b int) int {
	return arithmetic.Plus(a, b)
}

func Sun(a, b int) int {
	return arithmetic.Minus(a, b)
}
