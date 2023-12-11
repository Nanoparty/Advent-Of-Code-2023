package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
	//"unicode/utf8"
	//"strings"
	"time"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

type card_hand struct {
    hand string
    index  int
}

func main(){
	executionTime := time.Now()

	data, err := os.Open("test_input.txt")
	check(err)

	scanner := bufio.NewScanner(data)

	var bids map[int]int = make(map[int]int)

	five := []string[]

	line_num := 0
	for scanner.Scan() {
		line_num++
		line := scanner.Text()
		fmt.Println(line)

		hand := strings.Split(line, " ")[0]
		bid := strings.Split(line, " ")[1]

		bid_value, _ := strconv.Atoi(bid)
		bids[line_num] = bid_value



	}

	elapsed := time.Since(executionTime)
	fmt.Println("Done in", elapsed)
}