package bingoCard

import (
	"fmt"
	"strconv"
)

// Definition of slice of slice - slice of maps?
type BingoCard struct {
	Numbers     [][]string
	NumberMarks [][]bool
	Completed   bool
}

func (b *BingoCard) String() string {
	s := fmt.Sprintf("numbers: %s", b.Numbers)
	return s
}

func (b *BingoCard) MarkAllOccurances(markValue string) {
	for y, line := range b.Numbers {
		for x, number := range line {
			if number == markValue {
				b.NumberMarks[y][x] = true
			}
		}
	}
}

func (b *BingoCard) HasCardWon() bool {
	result := false
	// checking each row
	for _, line := range b.NumberMarks {
		if checkAllElementsTrue(line) == true {
			result = true
			//b.Completed = true
			return true
		}
	}
	// checking each column

	for x := range [5]int{} {
		var col []bool
		for y, _ := range b.NumberMarks {
			col = append(col, b.NumberMarks[y][x])
		}
		if checkAllElementsTrue(col) == true {
			result = true
			//b.Completed = true
			return true
		}
	}
	return result

}

func (b *BingoCard) FindUnmarkedTotal() int {
	sum := 0
	for y, line := range b.Numbers {
		for x, value := range line {
			if b.NumberMarks[y][x] == false {
				valueInt, _ := strconv.Atoi(value)
				sum += valueInt
			}
		}
	}
	return sum
}

// Constructor that builds object based off input of 5 lines
func NewBingoCard(input []string) BingoCard {
	newBingoCard := BingoCard{}
	newBingoCard.Numbers = append(newBingoCard.Numbers, input)
	newBingoCard.NumberMarks = append(newBingoCard.NumberMarks, []bool{false, false, false, false, false})
	return newBingoCard
}

func AddToBingoCard(card BingoCard, input []string) BingoCard {
	//cardLength := len(card.Numbers)
	card.Numbers = append(card.Numbers, input)
	card.NumberMarks = append(card.NumberMarks, []bool{false, false, false, false, false})
	return card
}

func checkAllElementsTrue(slice []bool) bool {
	result := true
	for _, element := range slice {
		if element != true {
			result = false
			break
		}
	}
	return result
}

// Mark all occurances of x

// CheckRow

// Checkcolumn
