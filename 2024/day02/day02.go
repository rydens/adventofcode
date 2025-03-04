package main

import (
	aoc "aoc2024/aochelpers"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var l [][]int

func main() {
	file := aoc.Noe(os.Open("./input.txt"))
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), " ")
		nums := aoc.Foreach(strs, aoc.Atoi)
		l = append(l, nums)
	}

	fmt.Println(part1())
	fmt.Println(part2())
}

func isSafe(is []int) bool {
	// increasing
	increasing := true
	for i := 0; i < len(is)-1; i++ {
		dif := is[i+1] - is[i]
		if (dif < 1) || (dif > 3) {
			increasing = false
		}
	}

	// decreasing
	decreasing := true
	for i := 0; i < len(is)-1; i++ {
		dif := is[i] - is[i+1]
		if (dif < 1) || (dif > 3) {
			decreasing = false
		}
	}

	return increasing || decreasing
}

func part1() int {
	return aoc.Count(l, isSafe)
}

func isSafe2(is []int) bool {
	if isSafe(is) {
		return true
	}

	try := 0
	for try < len(is) {
		newis := slices.Clone(is)
		newis = slices.Delete(newis, try, try+1)
		if isSafe(newis) {
			return true
		} else {
			try++
		}
	}

	return false
}

func part2() int {
	return aoc.Count(l, isSafe2)
}
