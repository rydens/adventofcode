package main

import (
	aoc "aoc2024/aochelpers"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var l1, l2 []int

func main() {
	file := aoc.Noe(os.Open("./input.txt"))
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), "   ")
		l1 = append(l1, aoc.Noe(strconv.Atoi(strs[0])))
		l2 = append(l2, aoc.Noe(strconv.Atoi(strs[1])))
	}

	slices.Sort(l1)
	slices.Sort(l2)

	part1()
	part2()
}

func part1() {
	sum := 0
	for i := 0; i < len(l1); i++ {
		sum += aoc.Abs(l1[i] - l2[i])
	}
	fmt.Println(sum)
}

func part2() {
	sum := 0
	for _, num := range l1 {
		sum += num * aoc.Count(l2, func(j int) bool { return j == num })
	}
	fmt.Println(sum)
}
