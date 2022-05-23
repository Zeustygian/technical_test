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
	if strings.Contains(arr_words, "euros/mois") == true {
		return true
	}
	if strings.Contains(arr_words, "EUR/mois") == true {
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
	if strings.Contains(arr_words, "EUR/an") == true {
		return true
	}
	if strings.Contains(arr_words, "euros/an") == true {
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
	if len(content) == 0 {
		fmt.Println("Empty string")
		os.Exit(84)
	}
	arr_words := strings.Split(content, " ")
	for i := 0; i != len(arr_words); i++ {
		if check_contain_year(arr_words[i]) == true {
			for prev_word := i; prev_word != 0; prev_word-- {
				price_year, err := strconv.Atoi(arr_words[prev_word])
				if price_year > 0 {
					fmt.Println("Price for a year is", price_year)
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
					fmt.Println("Price for a year is", price_year)
					return float32(price_year)
				}
				_ = err
			}
		}
	}
	return -1
}
