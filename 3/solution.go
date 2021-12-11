package main

import (
	"bufio"
	_ "errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("AoCinput3.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	horizontalInput := map[int][]string{}
	var inputList []string
	for scanner.Scan() {
		inputList = append(inputList, scanner.Text())
		seperatedLine := strings.Split(scanner.Text(), "")
		for idx, bit := range seperatedLine {
			//bitInt, _ := strconv.Atoi(bit)
			horizontalInput[idx] = append(horizontalInput[idx], bit)

		}

	}
	//fmt.Println(horizontalInput)
	//fmt.Println(horizontalInput)
	//gammaRate := ""
	//epsilonRate := ""

	keySlice := make([]int, 0)
	for key, _ := range horizontalInput {
		keySlice = append(keySlice, key)
	}

	// Now sort the slice
	sort.Ints(keySlice)

	var oxygenCandidates []string
	oxygenCandidates = append(oxygenCandidates, inputList...)
	var carbonCandidates []string
	carbonCandidates = append(carbonCandidates, inputList...)

	searchingForOxygen := true
	searchingForCarbon := true
	var oxygenDecimal int64
	var carbonDecimal int64

	for _, key := range keySlice {
		oxygenHorizontal := horizontaliseList(oxygenCandidates)
		carbonHorizontal := horizontaliseList(carbonCandidates)
		mostCommonBit := findMostCommonBit(oxygenHorizontal[key], "1")
		leastCommonBit := findLeastCommonBit(carbonHorizontal[key], "0")
		for _, value := range inputList {
			//mostCommonBitInt := mostCommonBit
			valueStr := string(value[key])
			if valueStr != mostCommonBit {
				oxygenCandidates = removeElement(oxygenCandidates, value)
			}
			if valueStr != leastCommonBit {
				carbonCandidates = removeElement(carbonCandidates, value)
			}

		}
		if len(oxygenCandidates) == 1 && searchingForOxygen {
			oxygenMeasure := oxygenCandidates[0]
			oxygenDecimal, _ = strconv.ParseInt(oxygenMeasure, 2, 64)
			searchingForOxygen = false
		}
		if len(carbonCandidates) == 1 && searchingForCarbon {
			carbonMeasure := carbonCandidates[0]
			carbonDecimal, _ = strconv.ParseInt(carbonMeasure, 2, 64)
			searchingForCarbon = false

		}
		if !searchingForOxygen && !searchingForCarbon {
			break
		}

	}
	answer := oxygenDecimal * carbonDecimal
	fmt.Println(answer)

	// 3366884 is too low

	//gammaRate += mostCommonBit
	//leastCommonBit := findLeastCommonBit(horizontalInput[key], preference "0")
	//epsilonRate += leastCommonBit
}

//fmt.Println("Gamma: " + gammaRate)
//fmt.Println("Epsilon: " + epsilonRate)
//gammaDecimal, _:= strconv.ParseInt(gammaRate, 2, 64)
//epsilonDecimal, _ := strconv.ParseInt(epsilonRate, 2, 64)
//
//fmt.Println("Gamma: ", gammaDecimal)
//fmt.Println("Epsilon: ",epsilonDecimal)
//result := gammaDecimal * epsilonDecimal
//
//fmt.Println(result)

// 4143194 too high

//}

func findMostCommonBit(bitList []string, preference string) string {
	zeroCount := 0
	oneCount := 0

	for _, bit := range bitList {
		if bit == "1" {
			oneCount++
		} else if bit == "0" {
			zeroCount++
		}
	}
	if oneCount > zeroCount {
		//fmt.Println("OneCount: ", oneCount, " ZeroCount: ", zeroCount)
		return "1"
	} else if oneCount < zeroCount {
		//fmt.Println("OneCount: ", oneCount, " ZeroCount: ", zeroCount)
		return "0"
	} else {
		return preference
	}
}

func findLeastCommonBit(bitList []string, preference string) string {
	zeroCount := 0
	oneCount := 0

	for _, bit := range bitList {
		if bit == "1" {
			oneCount++
		} else {
			zeroCount++
		}
	}
	if oneCount > zeroCount {
		return "0"
	} else if oneCount < zeroCount {
		return "1"
	} else {
		return preference
	}
}

func removeElement(slice []string, value string) []string {
	for idx, val := range slice {
		if val == value {
			return append(slice[:idx], slice[idx+1:]...)
		}
	}
	return slice
}

func horizontaliseList(slice []string) map[int][]string {
	horizontalMap := map[int][]string{}
	for _, value := range slice {
		seperatedLine := strings.Split(value, "")
		for idx, bit := range seperatedLine {
			//bitInt, _ := strconv.Atoi(bit)
			horizontalMap[idx] = append(horizontalMap[idx], bit)

		}
	}
	return horizontalMap
}
