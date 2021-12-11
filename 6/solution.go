package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("AoCInput6.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	//start, finish, axis := parseVentLineInput("3,3 -> 1,1")
	//grid.AddLine(start, finish, axis)
	//os.Exit(1)
	scanner := bufio.NewScanner(file)
	var fishCycles []string
	for scanner.Scan() {
		fishCycles = strings.Split(scanner.Text(), ",")
	}
	//fmt.Println(fishCycles)
	var fishCyclesMap = make(map[int]int)
	for _, day := range fishCycles {
		dayInt, _ := strconv.Atoi(day)
		_, found := fishCyclesMap[dayInt]
		if found {
			fishCyclesMap[dayInt]++
		} else {
			fishCyclesMap[dayInt] = 1
		}

	}

	days := 256
	for i := 0; i < days; i++ {
		temp := fishCyclesMap[0]
		for j := 0; j < 8; j++ {
			fishCyclesMap[j] = fishCyclesMap[j+1]
		}
		fishCyclesMap[8] = temp
		fishCyclesMap[6] += temp

	}

	count := 0

	for _, x := range fishCyclesMap {
		count += x
	}

	fmt.Println(count)

	//1895 is too low
	// 36405 is too low :(
}

//simulateDays(cycle)
//fmt.Println(fishCycles)
//var splitFishCycles [][]int
//size := 1
//var j int
//for i := 0; i < len(fishCyclesInt); i += size{
//	j += size
//	if j > len(fishCyclesInt) {
//		j = len(fishCyclesInt)
//	}
//	// do what do you want to with the sub-slice, here just printing the sub-slices
//	splitFishCycles = append(splitFishCycles, fishCyclesInt[i:j])
//}

//	fmt.Println(splitFishCycles)
//
//	cycleChan := make(chan []int)
//	count := 0
//	wg := &sync.WaitGroup{}
//
//	for _, cycle := range splitFishCycles{
//		wg.Add(1)
//		go simulateDays(cycle, 256, cycleChan, wg)
//
//	}
//	go func() {
//		wg.Wait()
//		close(cycleChan)
//	}()
//	//close(cycleChan)
//	for returnedCycle := range cycleChan{
//			count += len(returnedCycle)
//			//fmt.Println(count)
//
//	}
//	fmt.Println(count)
//
//	//for day, _ := range [18]int{} {
//	//	fmt.Println("Day: ", day)
//	//	var newCycleList []int
//	//	for _, cycle := range fishCyclesInt{
//	//		if cycle == 0{
//	//			newCycleList = append(newCycleList, 6) // reset counter on fish that spawned a new fish
//	//			newCycleList = append(newCycleList, 8) // add a new fish
//	//		}else{
//	//			newCycleList = append(newCycleList, cycle-1) // decrease fish towards day it will spawn new fish
//	//		}
//	//	}
//	//	fishCyclesInt = newCycleList
//	//	//fmt.Println(fishCyclesInt)
//	//}
//	//fmt.Println("FINAL")
//	//fmt.Println(fishCyclesInt)
//	//fmt.Println(len(fishCyclesInt))
//
//	//799284 is too high
//}
//
//
//func simulateDays(fishCycles []int, days int, cycleChan chan []int, wg *sync.WaitGroup)chan []int {
//	defer wg.Done()
//	for i:=0; i < days; i++ {
//		//fmt.Println("Day: ", i)
//		var newCycleList []int
//		for _, cycle := range fishCycles{
//			if cycle == 0{
//				newCycleList = append(newCycleList, 6) // reset counter on fish that spawned a new fish
//				newCycleList = append(newCycleList, 8) // add a new fish
//			}else{
//				newCycleList = append(newCycleList, cycle-1) // decrease fish towards day it will spawn new fish
//			}
//		}
//		fishCycles = newCycleList
//		//fmt.Println(fishCyclesInt)
//	}
//	cycleChan <- fishCycles
//	return cycleChan
//}
