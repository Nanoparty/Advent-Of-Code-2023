package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"time"
)

const RedCubes = 12
const GreenCubes = 13
const BlueCubes = 14

func check(e error){
	if e != nil {
		panic(e)
	}
}

func CheckValidGame(game string) int {
	gameId := strings.Split(strings.Split(game, ":")[0], " ")[1]

	setList := strings.Split(strings.Split(game, ":")[1],";")

	for _, set := range setList {
		
		picks := strings.Split(set, ",")
		for _, element := range picks {
			if (strings.Contains(element, "red")){
				num_string := strings.Split(element, " ")[1]
				num_val, err := strconv.Atoi(num_string)
				check(err)
				if num_val > RedCubes {
					fmt.Println("Invalid Game: ", gameId)
					return 0
				}
			}
			if (strings.Contains(element, "blue")){
				num_string := strings.Split(element, " ")[1]
				num_val, err := strconv.Atoi(num_string)
				check(err)
				if num_val > BlueCubes {
					fmt.Println("Invalid Game: ", gameId)
					return 0
				}
			}
			if (strings.Contains(element, "green")){
				num_string := strings.Split(element, " ")[1]
				num_val, err := strconv.Atoi(num_string)
				check(err)
				if num_val > GreenCubes {
					fmt.Println("Invalid Game: ", gameId)
					return 0
				}
			}
		}
	}
	
	game_value, err := strconv.Atoi(gameId)
	fmt.Printf("Valid Game: %d \n", game_value)
	check(err)
	return game_value
}

func CheckGamePower(game string) int {
	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	setList := strings.Split(strings.Split(game, ":")[1],";")

	for _, set := range setList {
		
		picks := strings.Split(set, ",")
		for _, element := range picks {
			if (strings.Contains(element, "red")){
				num_string := strings.Split(element, " ")[1]
				num_val, err := strconv.Atoi(num_string)
				check(err)
				if num_val > maxRed {
					maxRed = num_val
				}
			}
			if (strings.Contains(element, "blue")){
				num_string := strings.Split(element, " ")[1]
				num_val, err := strconv.Atoi(num_string)
				check(err)
				if num_val > maxBlue {
					maxBlue = num_val
				}
			}
			if (strings.Contains(element, "green")){
				num_string := strings.Split(element, " ")[1]
				num_val, err := strconv.Atoi(num_string)
				check(err)
				if num_val > maxGreen {
					maxGreen = num_val
				}
			}
		}
	}
	power := maxRed * maxBlue * maxGreen
	return power
}

func main() {
	executionTime := time.Now()

	sum := 0
	powerSum := 0

	data, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		sum += CheckValidGame(line)
		powerSum += CheckGamePower(line)
	}

	fmt.Printf("Sum: %d \n", sum)
	fmt.Printf("Power: %d \n", powerSum)

	elapsed := time.Since(executionTime)
	fmt.Println("Done in", elapsed)
}