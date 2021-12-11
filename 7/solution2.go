package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("AoCinput7.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	//start, finish, axis := parseVentLineInput("3,3 -> 1,1")
	//grid.AddLine(start, finish, axis)
	//os.Exit(1)
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = strings.Split(scanner.Text(), ",")
	}

	var inputInt []int
	for _, value := range input {
		valueInt, _ := strconv.Atoi(value)
		inputInt = append(inputInt, valueInt)
	}

	max := findLargestValue(inputInt)

	bestPos := 0
	bestCount := 1000000000

	for pos := 1; pos < max; pos++ {
		count := 0
		//cost := 1
		for _, value := range inputInt {
			difference := int(math.Abs(float64(value - pos)))
			for i := 1; i <= difference; i++ {
				count += i
			}
		}
		if count < bestCount {
			bestPos = pos
			bestCount = count
		}
	}
	fmt.Println(bestPos)
	fmt.Println(bestCount)

	//largestValue := findLargestValue(inputInt)
	//lowestValue := 0
	//fmt.Println(findCheapestPoint(inputInt, 80, lowestValue, largestValue))

	//353 is too low
	// 355 is too low haahahha
}

func findLargestValue(slice []int) int {
	var largerNumber, temp int
	for _, element := range slice {
		if element > temp {
			temp = element
			largerNumber = temp
		}
	}
	return largerNumber
}
