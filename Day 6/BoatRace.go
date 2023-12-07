package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"unicode/utf8"
	"strings"
	"time"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func CheckNumSolutions(time int, distance int) int {
	count := 0
	for i := 0; i <= time; i++ {
		remainingTime := time - i
		disTravelled := i * remainingTime
		if disTravelled > distance {
			count++
		}
	}
	return count
}

func main(){
	executionTime := time.Now()

	data, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(data)

	times := []int{}
	distances := []int{}

	big_time := ""
	big_distance := ""

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		if strings.Contains(line, "Time") {
			nums := strings.Split(strings.Split(line, ":")[1]," ")
			for _, element := range nums {
				num_val, err := strconv.Atoi(element)
				if err == nil {
					times = append(times, num_val)
					big_time += element
				}
			}
		}	

		if strings.Contains(line, "Distance") {
			nums := strings.Split(strings.Split(line, ":")[1]," ")
			for _, element := range nums {
				num_val, err := strconv.Atoi(element)
				if err == nil {
					distances = append(distances, num_val)
					big_distance += element
				}
			}
		}	
	}
	totalSolutions := 0

	for i := 0; i < len(times); i++ {
		result := CheckNumSolutions(times[i], distances[i])
		if totalSolutions == 0 {
			totalSolutions = result
		}else{
			totalSolutions *= result
		}
	}

	fmt.Println(times)
	fmt.Println(distances)
	fmt.Println("Solutions: ", totalSolutions)

	fmt.Println("Big Time: ", big_time)
	fmt.Println("Big Distance: ", big_distance)

	t,err := strconv.Atoi(big_time)
	d, err := strconv.Atoi(big_distance)

	big_solution := CheckNumSolutions(t, d)

	fmt.Println("Big Solution: ", big_solution)

	elapsed := time.Since(executionTime)
	fmt.Println("Done in", elapsed)
}