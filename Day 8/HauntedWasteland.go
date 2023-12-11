package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
	//"unicode/utf8"
	"strings"
	"time"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func main(){
	executionTime := time.Now()

	data, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(data)

	pattern := ""

	var leftSteps map[string]string = make(map[string]string)
	var rightSteps map[string]string = make(map[string]string)

	scanLine := 0

	for scanner.Scan() {
		scanLine++
		line := scanner.Text()
		//fmt.Println(line)

		if scanLine == 1 {
			pattern = line
			continue
		}
		if scanLine == 2 {
			continue
		}
		
		key := strings.Split(line, " = ")[0]
		values := strings.Split(line, " = ")[1]

		value1 := strings.Split(values, ", ")[0][1:]
		value2 := strings.Split(values, ", ")[1]
		value2 = value2[:len(value2)-1]

		leftSteps[key] = value1
		rightSteps[key] = value2
	}

	steps := 0
	node := "AAA"
	var dir int = 0
	for node != "ZZZ" {
		steps++
		direction := string(pattern[dir])
		if direction == "R" {
			node = rightSteps[node]
		}
		if direction == "L" {
			node = leftSteps[node]
		}
		dir++
		if dir >= len(pattern) {
			dir = 0
		}
	}	

	fmt.Println("TOTAL STEPS:", steps)

	elapsed := time.Since(executionTime)
	fmt.Println("Done in", elapsed)
}