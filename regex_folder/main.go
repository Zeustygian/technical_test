package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check_contain_month(arr_words string) bool {
	if strings.Contains(arr_words, "euro/mois") == true {
		return true
	}
	if strings.Contains(arr_words, "/mois") == true {
		return true
	}
	if strings.Contains(arr_words, "mois") == true {
		return true
	}
	return false
}

func check_contain_year(arr_words string) bool {
	if strings.Contains(arr_words, "euro/an") == true {
		return true
	}
	if strings.Contains(arr_words, "/an") == true {
		return true
	}
	if strings.Contains(arr_words, "an") == true {
		return true
	}
	return false
}

func ExtractRent(content string) float32 {
	arr_words := strings.Split(content, " ")
	for i := 0; i != len(arr_words); i++ {
		if check_contain_year(arr_words[i]) == true {
			for prev_word := i; prev_word != 0; prev_word-- {
				price_year, err := strconv.Atoi(arr_words[prev_word])
				if price_year > 0 {
					fmt.Println(price_year)
					return float32(price_year)
				}
				_ = err
			}
		}
		if check_contain_month(arr_words[i]) == true {
			for prev_word := i; prev_word != 0; prev_word-- {
				price_year, err := strconv.Atoi(arr_words[prev_word])
				if price_year > 0 {
					price_year *= 12
					fmt.Println(price_year)
					return float32(price_year)
				}
				_ = err
			}
		}
	}
	return 84
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(84)
	}
	content := os.Args[1]
	if ExtractRent(content) == 84 {
		os.Exit(84)
	} else {
		os.Exit(0)
	}
}
