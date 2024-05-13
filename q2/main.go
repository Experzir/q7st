package main

import (
	"fmt"
)

func main() {
	// var input string
	// fmt.Println("Enter your input :")
	// fmt.Scanln(&input)
	// fmt.Printf("Input is : %v\n",input)
	// fmt.Println(decode(input))

	test := []string{"LLRR=", "==RLL", "=LLRR", "RRL=R"}
	for _, v := range test {
		res := decode(v)
		fmt.Printf("input = %v\toutput = %v\n", v, res)
	}
}

func decode(s string) []int {
	// result := ""
	res := make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == 'L' {
				res[i] = 1
				res[j] = 0
			} else if s[i] == 'R' {
				res[i] = 0
				res[j] = 1
			} else if s[i] == '=' {
				if i != 0 {
					res[j] = res[i]
				} else {
					res[j] = 0
					res[i] = res[j]
				}
			} else {
				return []int{0}
			}
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			if !(res[i] > res[i+1]) {
				res[i] = res[i+1] + 1
			}
		} else if s[i] == 'R' {
			if !(res[i+1] > res[i]) {
				res[i+1] = res[i] + 1
			}
		} else if s[i] == '=' {
			res[i+1] = res[i]
		}
	}
	for i := 1; i < len(s); i++ {
		if s[i-1] == 'L' {
			if !(res[i-1] > res[i]) {
				res[i-1] = res[i] + 1
			}
		} else if s[i-1] == 'R' {
			if !(res[i] > res[i-1]) {
				res[i] = res[i-1] + 1
			}
		} else if s[i-1] == '=' {
			res[i-1] = res[i]
		}
	}
	return res
}
