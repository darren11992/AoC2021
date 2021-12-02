package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	file, err := os.Open("AoCinput1.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	increaseCount := 0
	decreaseCount := 0
	noChangeCount := 0
	lineNo := 0
	prevSlidingWindowSum := 0
	var slidingWindow []int
	//currentMeasurement := 0
	//var currentMeasurement int
	for scanner.Scan() {
		if len(slidingWindow) == 3{
			newSum := sum(slidingWindow)
			if newSum < prevSlidingWindowSum && lineNo != 3{
				decreaseCount++
			}else if newSum > prevSlidingWindowSum && lineNo != 3{
				increaseCount++
			}else if newSum == prevSlidingWindowSum && lineNo != 3{
				noChangeCount++
			}

			newMeasurement, _ := strconv.Atoi(scanner.Text())
			slidingWindow = append(slidingWindow, newMeasurement)
			slidingWindow = slidingWindow[1:]

			prevSlidingWindowSum = newSum
			lineNo++
		}else{
			newMeasurement, _ := strconv.Atoi(scanner.Text())
			slidingWindow = append(slidingWindow, newMeasurement)
			//fmt.Println(slidingWindow)
			lineNo++
		}
		//fmt.Println(scanner.Text())
	}
	fmt.Println("inc", increaseCount)
	fmt.Println("dec", decreaseCount)
	fmt.Println("No", noChangeCount)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}


}

func sum(s []int) int{
	total := 0
	for _, value := range s{
		total += value
	}
	return total
}
