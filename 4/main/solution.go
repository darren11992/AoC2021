package main

import (
	. "AoC2021/4/bingoCard"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("AoCinput4.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	//scanner.Split(bufio.ScanLines)
	var lines []string
	var bingoDraw string
	for scanner.Scan() {
		//bingoDraw = scanner.Text()
		//fmt.Println(bingoDraw)
		if scanner.Text() != "" {
			lines = append(lines, scanner.Text())
		}

		// putting the file into a slice so i can access indexs
	}
	bingoDraw = lines[0]
	//file.Close()

	var bingoCards []BingoCard

	for idx, line := range lines[1:] {
		//fmt.Println("Line ", idx, " : ", line)
		stripedLine := strings.TrimSpace(line)
		splitLine := strings.Split(stripedLine, " ")
		splitLine = removeMiddleWhiteSpace(splitLine)
		if idx%5 == 0 {
			newCard := NewBingoCard(splitLine)
			bingoCards = append(bingoCards, newCard)
		} else {
			bingoCards[idx/5] = AddToBingoCard(bingoCards[idx/5], splitLine)
		}

	}

	bingoDrawSplit := strings.Split(bingoDraw, ",")
	for _, markValue := range bingoDrawSplit {
		fmt.Println(markValue)
		for idx, card := range bingoCards {
			card.MarkAllOccurances(markValue)
			if card.HasCardWon() {
				fmt.Println("Card ", idx, "has BINGO!")
				bingoCards[idx].Completed = true
				cardValue := card.FindUnmarkedTotal()
				markValueInt, _ := strconv.Atoi(markValue)
				result := cardValue * markValueInt
				fmt.Println("result: ", result)
			}
		}
		bingoCards = removeFinishedBingoCards(bingoCards)
	}

	//for idx, card := range bingoCards{
	//	if card.HasCardWon(){
	//		fmt.Println("Card ", idx, "has BINGO!")
	//	}
	//}

}

func removeMiddleWhiteSpace(line []string) []string {
	var newLine []string
	for _, value := range line {
		if value != "" {
			newLine = append(newLine, value)
		}
	}
	return newLine
}

func removeFinishedBingoCards(bingoCards []BingoCard) []BingoCard {
	var newCardSlice []BingoCard
	for _, card := range bingoCards {
		if card.Completed == false {
			newCardSlice = append(newCardSlice, card)
		}
	}
	return newCardSlice
}
