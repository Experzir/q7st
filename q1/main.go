package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// {
// 	     {59},
// 	   {73, 41},
// 	 {52, 40, 53},
// 	{26, 53, 6, 34},
// }

func main() {
	var arr [][]int

	jsonFile, err := os.Open("hard.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(&arr)
	if err != nil {
		log.Fatal(err)
	}

	check := make(map[[2]int]int)

	fmt.Println(maxPath(arr,0,0,check))
}

func maxPath(arr [][]int, row, col int, check map[[2]int]int) int {
	// already value
	if val, ok := check[[2]int{row, col}]; ok {
		return val
	}

	if row == len(arr)-1 {
		return arr[row][col]
	}

	left := maxPath(arr, row+1, col, check)
	right := maxPath(arr, row+1, col+1, check)

	maxValue := arr[row][col] + maxNum(left, right)

	// set value
	check[[2]int{row, col}] = maxValue

	return maxValue
}

func maxNum(x, y int) int {
	if x > y {
		return x
	}
	return y
}
