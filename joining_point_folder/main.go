package main

import (
	"fmt"
	"os"
	"strconv"
)

func Compute_first_nb(first_string string, sum_first int) int {
	for i := 0; i < len(first_string); i++ {
		char := string(first_string[i])
		value, err := strconv.Atoi(char)
		_ = err
		sum_first += value
	}
	return sum_first
}

func Compute_snd_nb(snd_string string, sum_snd int) int {
	for i := 0; i != len(snd_string); i++ {
		char := string(snd_string[i])
		value, err := strconv.Atoi(char)
		_ = err
		sum_snd += value
	}
	return sum_snd
}

func ComputeJoinPoint(s1 int, s2 int) int {
	if s1 >= 20000000 || s2 >= 20000000 || s1 <= 0 || s2 <= 0 {
		fmt.Println("Invalid parameter value.")
		os.Exit(84)
	}
	sum_first := s1
	sum_snd := s2
	first_turn := false
	for sum_first != sum_snd || first_turn != true {
		cpy_first_str := strconv.Itoa(sum_first)
		cpy_snd_str := strconv.Itoa(sum_snd)
		sum_first = Compute_first_nb(cpy_first_str, sum_first)
		sum_snd = Compute_snd_nb(cpy_snd_str, sum_snd)
		first_turn = true
		if sum_first < 0 || sum_first >= 20000000 {
			fmt.Println("Error :The join point exceed 20000000")
			os.Exit(84)
		}
	}
	if sum_first >= 0 || sum_first < 20000000 {
		fmt.Println("The joining point of these numbers is", sum_first)
		os.Exit(0)
	}
	return sum_first
}
