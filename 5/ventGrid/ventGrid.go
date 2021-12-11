package ventGrid

import (
	"fmt"
)

type VentGrid struct {
	points [][]int
}

func NewVentGrid(length int, width int) VentGrid {
	newGrid := make([][]int, length)
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			newGrid[i] = append(newGrid[i], 0)
		}
	}
	return VentGrid{newGrid}

}

func (v *VentGrid) Stringy() {
	for idx, _ := range v.points {
		fmt.Println(v.points[idx])

	}
}

func (v *VentGrid) AddLine(start []int, finish []int, lineAxis string) {
	if lineAxis == "x" {
		yValue := start[1]
		for x := start[0]; x <= finish[0]; x++ {
			v.points[yValue][x]++
		}
	} else if lineAxis == "y" {
		xValue := start[0]
		for y := start[1]; y <= finish[1]; y++ {
			v.points[y][xValue]++
		}
	} else {
		// diagonal - ordered by x, so will either be /^ or \v directions
		fmt.Println("diag")
		y := start[1]
		if finish[1] > start[1] {
			//right-up diagonal

			for x := start[0]; x <= finish[0]; x++ {
				v.points[y][x]++
				y++
				//for y := start[1]; y<=finish[1]; y++{
				//	v.points[y][x]++
				//}
			}
		} else {
			//right-down diagonal
			for x := start[0]; x <= finish[0]; x++ {
				v.points[y][x]++
				y--

				//for y := start[1]; y>=finish[1]; y--{
				//	v.points[y][x]++
				//}
			}
		}
	}

}

func (v *VentGrid) GetOverlapCount() int {
	count := 0
	for _, y := range v.points {
		for _, value := range y {
			if value > 1 {
				count++
			}
		}
	}
	return count
}
