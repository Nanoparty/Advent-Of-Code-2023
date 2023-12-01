package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
	"strings"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func CheckForNumber(substring string) string {
	if strings.Contains(substring, "one") {
		return "1"
	}
	if strings.Contains(substring, "two") {
		return "2"
	}
	if strings.Contains(substring, "three") {
		return "3"
	}
	if strings.Contains(substring, "four") {
		return "4"
	}
	if strings.Contains(substring, "five") {
		return "5"
	}
	if strings.Contains(substring, "six") {
		return "6"
	}
	if strings.Contains(substring, "seven") {
		return "7"
	}
	if strings.Contains(substring, "eight") {
		return "8"
	}
	if strings.Contains(substring, "nine") {
		return "9"
	}
	return "-1"
}

func FirstNumber(input string) string {
	index := 0
	for _,c := range input {

		if c >= '0' && c <= '9' {
			return string(c)
		} else {
			var substring string = input[0:index+1]
			NumberCheck := CheckForNumber(substring)
			if NumberCheck != "-1" {
				return NumberCheck
			}
		}
		index++
	}
	return ""
}

func LastNumber(input string) string {
	inputLength := utf8.RuneCountInString(input)
	for i := inputLength-1; i >= 0; i-- {
		runeValue, _ := utf8.DecodeRuneInString(input[i:i+1])
		if runeValue >= '0' && runeValue <= '9' {
			return string(runeValue)
		} else {
			var substring string = input[i:inputLength]
			NumberCheck := CheckForNumber(substring)
			if NumberCheck != "-1" {
				return NumberCheck
			}
		}
	}
	return ""
}

func main() {

	sum := 0

	dat, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		line := scanner.Text()

		firstNumber := FirstNumber(line)

		lastNumber := LastNumber(line)

		combinedNumber := firstNumber + lastNumber

		lineValue, err := strconv.Atoi(combinedNumber)
		
		if err != nil {
			panic(err)
		}

		sum += lineValue
	}

	dat.Close()

	fmt.Printf("SUM: %d", sum)
}