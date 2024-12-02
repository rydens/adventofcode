package aochelpers

func Noe[A any](a A, err error) A {
	if err != nil {
		panic(err)
	}
	return a
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
