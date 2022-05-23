package main

import (
	"fmt"
	"os"
)

func Calc(arr []int, n1, n2 int) int {
	value := 0
	for i := 0; i < len(arr); i++ {
		if n1 < 0 || n2 > len(arr)-1 {
			os.Exit(84)
		}
		if i >= n1 && i <= n2 {
			value += arr[i]
		}
	}
	fmt.Printf("The sum of numbers between argument[%d] and argument[%d] is %d\n", n1, n2, value)
	return value
}

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Don't give arguments to this binary please.")
		os.Exit(84)
	}
	arr_sliced := []int{0, 1, 2, 3}
	if len(arr_sliced) == 0 {
		fmt.Println("Empty array.")
		os.Exit(84)
	}
	n1 := 3
	n2 := 0
	if n1 > n2 {
		fmt.Println("n1 variabe can't be superior than n2 variable.")
		os.Exit(84)
	}
	fmt.Println(arr_sliced)
	Calc(arr_sliced[:], n1, n2)
}
