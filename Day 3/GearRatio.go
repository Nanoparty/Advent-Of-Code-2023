package main

import (
	"fmt"
	"os"
	"bufio"
	"time"
	"unicode/utf8"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func main() {
	executionTime := time.Now()

	data, err := os.Open("small_input.txt")
	check(err)

	scanner := bufio.NewScanner(data)

	var rows int = 0
	var cols int = 0
	for scanner.Scan() {
		line := scanner.Text()
		cols = utf8.RuneCountInString(line)
		rows++
	}

	fmt.Println("Rows:", rows, " Cols:", cols)

	schematic := make([][]string, rows)
	for i := range schematic {
		schematic[i] = make([]string, cols)
	}

	schematic[0][0] = "hi"

	fmt.Println(schematic[0][0])

	scanner2 := bufio.NewScanner(data)

	r := 0
	for scanner2.Scan() {
		line := scanner2.Text()
		fmt.Print(line)
		for i := 0; i < utf8.RuneCountInString(line); i++{
			runeValue, _ := utf8.DecodeRuneInString(line[i:i+1])
			schematic[r][i] = string(runeValue)
			fmt.Print(schematic[r][i])
		}
		r++
	}

	// for r := 0; r < rows;r++ {
	// 	for c := 0; c < cols;c++ {
	// 		fmt.Print(schematic[r][c], " ")
	// 	}
	// 	fmt.Print("\n")
	// }

	elapsed := time.Since(executionTime)
	fmt.Println("Done in", elapsed)
}