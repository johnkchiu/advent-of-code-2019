package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var totalFuel = []int{}

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input, _ := strconv.ParseFloat(scanner.Text(), 64)
		fuel := getFuel(input)
		fmt.Printf("%v == %v\n", input, fuel)
		totalFuel = append(totalFuel, fuel)
	}

	var total = 0
	for _, value := range totalFuel {
		total += value
	}
	fmt.Printf("total: %v", total)
}

func getFuel(mass float64) int {
	return int(math.Trunc(mass/3) - 2)
}
