package main

import (
	"AoC2021/5/ventGrid"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	grid := ventGrid.NewVentGrid(1000, 1000)

	file, err := os.Open("AoCInput5.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	//start, finish, axis := parseVentLineInput("3,3 -> 1,1")
	//grid.AddLine(start, finish, axis)
	//os.Exit(1)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		start, finish, axis := parseVentLineInput(scanner.Text())
		grid.AddLine(start, finish, axis)

	}

	grid.Stringy()
	fmt.Println(grid.GetOverlapCount())
}

func parseVentLineInput(line string) ([]int, []int, string) {
	trimLine := strings.TrimSpace(line)
	splitLine := strings.Split(trimLine, "->")
	//splitLine = removeMiddleWhiteSpace(splitLine)
	firstCoordinate := strings.Split(splitLine[0], ",")
	secondCoordinate := strings.Split(splitLine[1], ",")
	intFirstCoordinate := stringSliceToInt(firstCoordinate)
	intSecondCoordinate := stringSliceToInt(secondCoordinate)

	if intFirstCoordinate[0] == intSecondCoordinate[0] {
		// x coordinates the same- "vertical line"
		lineAxis := "y"
		if intFirstCoordinate[1] < intSecondCoordinate[1] {
			return intFirstCoordinate, intSecondCoordinate, lineAxis
		} else {
			return intSecondCoordinate, intFirstCoordinate, lineAxis
		}
	} else if intFirstCoordinate[1] == intSecondCoordinate[1] {
		// y coordinates the same- "horizontal line"
		lineAxis := "x"
		if intFirstCoordinate[0] < intSecondCoordinate[0] {
			return intFirstCoordinate, intSecondCoordinate, lineAxis
		} else {
			return intSecondCoordinate, intFirstCoordinate, lineAxis
		}
	} else {
		//Invalid "line" from given coordinates (in part 1 haha)
		// Now a "diagonal line"
		// Order by the x axis- let
		fmt.Println("Coordinates given do not represent a straight line")
		lineAxis := "diagonal"
		if intFirstCoordinate[0] < intSecondCoordinate[0] {

			return intFirstCoordinate, intSecondCoordinate, lineAxis
		} else {
			return intSecondCoordinate, intFirstCoordinate, lineAxis
		}
	}
}

func stringSliceToInt(stringSlice []string) []int {
	var intSlice []int
	for _, value := range stringSlice {
		strippedValue := strings.TrimSpace(value)
		valueInt, _ := strconv.Atoi(strippedValue)
		intSlice = append(intSlice, valueInt)
	}
	return intSlice
}
