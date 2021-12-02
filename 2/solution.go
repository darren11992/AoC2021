package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	file, err := os.Open("AoCinput2.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	const forwardCmd = "forward"
	const downCmd = "down" // add to depth
	const upCmd = "up" // subtract from depth

	depth := 0
	horizontalPosition := 0
	aim := 0

	for scanner.Scan() {
		//seperatedLine := scanner.Text()
		seperatedLine := strings.Split(scanner.Text(), " ")
		command := seperatedLine[0]
		commandAmount, _ := strconv.Atoi(seperatedLine[1])

		if command == forwardCmd{
			horizontalPosition += commandAmount
			depth += (commandAmount * aim)
		}

		if command == upCmd{
			aim -= commandAmount
		}

		if command == downCmd{
			aim += commandAmount
		}


	}

	result := depth * horizontalPosition
	fmt.Println(result)
}