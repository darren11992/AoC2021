package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("testInput.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string
	var count int
	for scanner.Scan() {
		input = strings.Split(scanner.Text(), "|")
		inputValue := strings.Split(input[0], " ")

		eightValue := findEightString(inputValue)

		//if eightValue == "Not Found"{
		//	fmt.Println("Eightvalue not found :(")
		//	break
		//}

		encodings := getEncoding(eightValue)

		outputValue := strings.Split(input[1], " ")
		var thisLinesNumberString string
		for _, word := range outputValue {
			thisLinesNumberString += stringToEncodingNo(word, encodings)

			//part 1
			//wordStripped := strings.TrimSpace(word)
			//if (len(wordStripped) >= 2 && len(wordStripped) <= 4) || len(wordStripped) == 7 {
			//	count++
			//}
		}
		thisLinesNumber, _ := strconv.Atoi(thisLinesNumberString)
		count += thisLinesNumber

	}
	fmt.Println(count)

	// 19836 is too low

	getEncoding("acedgfb")

}

//digits 1, 4, 7, and 8 use 2, 4, 3 and 7 segments respectivly
//map[int]string
func getEncoding(eightString string) map[string]string {
	eightString = sortString(eightString)
	eightSplit := strings.Split(eightString, "")
	//TODO: This can be fixed by ordering alphabetically the "eightvalue" you have on your ipad, as well
	// as the other values you have, and then finding a index pattern in the same way as you have done before,
	// then update this function :)
	oneString := sortString(eightSplit[0] + eightSplit[1])
	twoString := sortString(eightSplit[0] + eightSplit[2] + eightSplit[3] + eightSplit[5] + eightSplit[6])
	threeString := sortString(eightSplit[0] + eightSplit[1] + eightSplit[2] + eightSplit[3] + eightSplit[5])
	fourString := sortString(eightSplit[0] + eightSplit[1] + eightSplit[4] + eightSplit[5])
	fiveString := sortString(eightSplit[1] + eightSplit[2] + eightSplit[3] + eightSplit[4] + eightSplit[5])
	sixString := sortString(eightSplit[1] + eightSplit[2] + eightSplit[3] + eightSplit[4] + eightSplit[5] + eightSplit[6])
	sevenString := sortString(eightSplit[0] + eightSplit[1] + eightSplit[3])
	sortedEightString := sortString(eightSplit[0] + eightSplit[1] + eightSplit[2] + eightSplit[3] + eightSplit[4] + eightSplit[5] + eightSplit[6])
	nineString := sortString(eightSplit[0] + eightSplit[1] + eightSplit[2] + eightSplit[3] + eightSplit[4] + eightSplit[5])
	zeroString := sortString(eightSplit[0] + eightSplit[1] + eightSplit[2] + eightSplit[3] + eightSplit[4] + eightSplit[6])

	returnList := map[string]string{zeroString: "0", oneString: "1", twoString: "2", threeString: "3", fourString: "4", fiveString: "5", sixString: "6", sevenString: "7", sortedEightString: "8", nineString: "9"}
	return returnList
}

func findEightString(strings []string) string {
	for _, i := range strings {
		if len(i) == 7 {
			return i
		}
	}
	return "Not Found"
}

func stringToEncodingNo(str string, encodings map[string]string) string {

	sortedString := sortString(str)

	numberAsString, found := encodings[sortedString]
	if found {
		return numberAsString
	} else {
		return ""
	}
}

func sortString(word string) string {
	s := []rune(word)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	return string(s)
}
