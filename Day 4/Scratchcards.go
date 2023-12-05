package main

import (
	"fmt"
	"os"
	"bufio"
	"time"
	"strings"
	"strconv"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func CheckCard(input string) int {
	all_numbers := strings.Split(input, ":")[1]
	winning_numbers := strings.Split(all_numbers, "|")[0]
	my_numbers := strings.Split(all_numbers, "|")[1]

	my_list := strings.Split(my_numbers, " ")
	winning_list := strings.Split(winning_numbers, " ")

	my_slice := []int{}
	for _, element := range my_list {
		num_val, err := strconv.Atoi(element)
		if err == nil {
			my_slice = append(my_slice, num_val)
		}
	}

	winning_slice := []int{}
	for _, element := range winning_list {
		num_val, err := strconv.Atoi(element)
		if err == nil {
			winning_slice = append(winning_slice, num_val)
		}
	}

	total_points := 0
	for _, my_element := range my_slice {
		for _, winning_element := range winning_slice {
			if my_element == winning_element {
				if total_points == 0 {
					total_points = 1
				}else {
					total_points = total_points * 2
				}
			}
		}
	}
	return total_points
}

func main(){
	executionTime := time.Now()

	dat, err := os.Open("input.txt")
	check(err)

	total := 0

	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		line := scanner.Text()		
		total += CheckCard(line)
	}

	dat.Close()

	fmt.Println("Total Points: ", total)

	elapsed := time.Since(executionTime)
	fmt.Println("Done in", elapsed)
}

