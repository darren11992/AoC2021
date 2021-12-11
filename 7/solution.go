package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func oldmain() {
	file, err := os.Open("testInput.txt")

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

	largestValue := findLargestValue(inputInt)
	lowestValue := 0
	fmt.Println(findCheapestPoint(inputInt, 80, lowestValue, largestValue))

	//353 is too low
	// 355 is too low haahahha

	// 999 is too low
}

func findCheapestPoint(input []int, increment int, lowestValue int, maxValue int) int {
	fmt.Println("new call! LowestValue: ", lowestValue, " highest value: ", maxValue, "increment: ", increment)
	if maxValue-lowestValue == 1 {
		return lowestValue
	} else {
		lowestCost := 1000000
		var bestIteration int
		var secondBestIteration int
		secondBestCost := 1000000
		var costs = make(map[int]int)
		var costsSlice []int
		for i := lowestValue; i <= maxValue; i += increment {
			currentCost := 0
			for _, value := range input {
				currentCost += int(math.Abs(float64(value - i)))
			}
			costs[i] = currentCost
			costsSlice = append(costsSlice, currentCost)
			//fmt.Println("Current iteration: ", i, " Cost: ", currentCost)
			if currentCost < lowestCost {
				lowestCost = currentCost
				bestIteration = i
			}
			currentCost = 0

		}

		for key, cost := range costs {
			if cost < secondBestCost && cost > lowestCost {
				secondBestCost = cost
				secondBestIteration = key
			}
		}
		//sort.Ints(costs)
		//fmt.Println("Lowest: ", bestIteration, "with cost: ", lowestCost)
		//fmt.Println("2nd best: ", secondBestIteration, "with cost: ", secondBestCost)
		nextIncrement := findBestIncrement(increment/2, bestIteration, secondBestIteration)

		if bestIteration > secondBestIteration {
			return findCheapestPoint(input, nextIncrement, secondBestIteration, bestIteration)
		} else {
			return findCheapestPoint(input, nextIncrement, bestIteration, secondBestIteration)
		}
	}
}

//func findLargestValue(slice []int)int{
//	var largerNumber, temp int
//	for _, element := range slice {
//		if element > temp {
//			temp = element
//			largerNumber = temp
//		}
//	}
//	return largerNumber
//}

func findBestIncrement(maxIterations int, lowValue int, highValue int) int {
	totalValues := highValue - lowValue
	if totalValues < maxIterations {
		return 1
	} else {
		return totalValues / maxIterations
	}
}
