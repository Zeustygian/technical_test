package main

import (
	"fmt"
	"os"
)

func Calc(arr []int, n1, n2 int) int {
	if n1 > n2 {
		fmt.Println("n1 variabe can't be superior than n2 variable.")
		os.Exit(84)
	}
	if len(arr) == 0 {
		fmt.Println("Empty array.")
		os.Exit(84)
	}
	value := 0
	for i := 0; i < len(arr); i++ {
		if n1 < 0 || n2 > len(arr)-1 {
			fmt.Println("n1 or n2 is out of the array's limits")
			os.Exit(84)
		}
		if i >= n1 && i <= n2 {
			value += arr[i]
		}
	}
	fmt.Printf("The sum of numbers between index[%d] and index[%d] is %d\n", n1, n2, value)
	return value
}
