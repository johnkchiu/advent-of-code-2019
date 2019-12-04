package main

import (
	"fmt"

	"github.com/johnkchiu/adventofcode2019/common"
)

func main() {
	test("1,9,10,3,2,3,11,0,99,30,40,50")
	test("1,0,0,0,99")
	test("2,3,0,3,99")
	test("2,4,4,5,99,0")
	test("1,1,1,4,99,5,6,0,99")

	// test("1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,10,19,23,1,23,9,27,1,5,27,31,2,31,13,35,1,35,5,39,1,39,5,43,2,13,43,47,2,47,10,51,1,51,6,55,2,55,9,59,1,59,5,63,1,63,13,67,2,67,6,71,1,71,5,75,1,75,5,79,1,79,9,83,1,10,83,87,1,87,10,91,1,91,9,95,1,10,95,99,1,10,99,103,2,103,10,107,1,107,9,111,2,6,111,115,1,5,115,119,2,119,13,123,1,6,123,127,2,9,127,131,1,131,5,135,1,135,13,139,1,139,10,143,1,2,143,147,1,147,10,0,99,2,0,14,0")
	day2("1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,10,19,23,1,23,9,27,1,5,27,31,2,31,13,35,1,35,5,39,1,39,5,43,2,13,43,47,2,47,10,51,1,51,6,55,2,55,9,59,1,59,5,63,1,63,13,67,2,67,6,71,1,71,5,75,1,75,5,79,1,79,9,83,1,10,83,87,1,87,10,91,1,91,9,95,1,10,95,99,1,10,99,103,2,103,10,107,1,107,9,111,2,6,111,115,1,5,115,119,2,119,13,123,1,6,123,127,2,9,127,131,1,131,5,135,1,135,13,139,1,139,10,143,1,2,143,147,1,147,10,0,99,2,0,14,0")
}

func test(str string) {
	ints, _ := common.Split(str, ",")
	intcode(ints)
}

func day2(str string) {
	ints, _ := common.Split(str, ",")
	fmt.Printf("og: %v\n", str)
	ints[1] = 12
	ints[2] = 2
	fmt.Printf("mod: %v\n", str)
	intcode(ints)
}

func intcode(prog []int) {
	fmt.Printf("before: %v\n", prog)

	current := 0
	op := 0
	for ; op < 99; op = getOp(current, prog) {
		// fmt.Printf("op: %v\n", op)
		switch op {
		case 1:
			current = execAdd(current, prog)
		case 2:
			current = execMultiply(current, prog)
		default:
			// no op
		}
	}

	fmt.Printf("after: %v\n", prog)
}

func getOp(c int, prog []int) int {
	return prog[c]
}

func execAdd(c int, prog []int) int {
	first := prog[prog[c+1]]
	second := prog[prog[c+2]]
	prog[prog[c+3]] = first + second
	// fmt.Printf("prog[%v] = %v + %v\n", c+3, first, second)
	return c + 4
}

func execMultiply(c int, prog []int) int {
	first := prog[prog[c+1]]
	second := prog[prog[c+2]]
	prog[prog[c+3]] = first * second
	// fmt.Printf("prog[%v] = %v * %v\n", c+3, first, second)
	return c + 4
}
