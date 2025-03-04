package aochelpers

import (
	"strconv"
)

func Noe[A any](a A, err error) A {
	if err != nil {
		panic(err)
	}
	return a
}

func Atoi(s string) int {
	return Noe(strconv.Atoi(s))
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Count[T any](slice []T, f func(T) bool) int {
	count := 0
	for _, elem := range slice {
		if f(elem) {
			count++
		}
	}
	return count
}

func Foreach[T any, A any](slice []T, f func(T) A) []A {
	a := make([]A, len(slice))
	for i, elem := range slice {
		a[i] = f(elem)
	}
	return a
}
