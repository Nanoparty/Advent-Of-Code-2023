package main

import (
	"fmt"
	"os"
	"bufio"
	"time"
	"strings"
	"strconv"
	//"slices"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func CreateMap(line string, input map[uint64]uint64) {
	soil_list := strings.Split(line, " ")
	dest_start, _ := strconv.Atoi(string(soil_list[0]))
	source_start, _ := strconv.Atoi(string(soil_list[1]))
	range_len, _ := strconv.Atoi(string(soil_list[2]))

	dest_start64 := int64(dest_start)
	source_start64 := int64(source_start)
	range_len64 := int64(range_len)
	
	for i := int64(0); i < range_len64; i++ {
		input[source_start64 + i] = dest_start64 + i;
	}
}

func main(){
	executionTime := time.Now()

	dat, err := os.Open("input.txt")
	check(err)

	seeds := []uint64{}

	var soil map[uint64]uint64 = make(map[uint64]uint64)
	var fertilizer map[uint64]uint64 = make(map[uint64]uint64)
	var water map[uint64]uint64 = make(map[uint64]uint64)
	var light map[uint64]uint64 = make(map[uint64]uint64)
	var temperature map[uint64]uint64 = make(map[uint64]uint64)
	var humidity map[uint64]uint64 = make(map[uint64]uint64)
	var location map[uint64]uint64 = make(map[uint64]uint64)

	var reader string = ""
	line_num := 1

	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		fmt.Println("Line #", line_num)
		line_num++
		line := scanner.Text()	

		if line == ""{
			reader = ""
			continue;
		}
		if strings.Contains(line, "seed-to-soil") {
			reader = "soil"
			continue
		}
		if strings.Contains(line, "soil-to-fertilizer") {
			reader = "fertilizer"
			continue
		}
		if strings.Contains(line, "fertilizer-to-water") {
			reader = "water"
			continue
		}
		if strings.Contains(line, "water-to-light") {
			reader = "light"
			continue
		}
		if strings.Contains(line, "light-to-temperature") {
			reader = "temperature"
			continue
		}
		if strings.Contains(line, "temperature-to-humidity") {
			reader = "humidity"
			continue
		}
		if strings.Contains(line, "humidity-to-location") {
			reader = "location"
			continue
		}

		// Read Seeds
		if strings.Contains(line, "seeds:")	{
			seeds_string := strings.Split(line, ":")[1]
			seeds_list := strings.Split(seeds_string, " ")
			for _, element := range seeds_list {
				num_val, err := strconv.Atoi(element)
				if err == nil {
					seeds = append(seeds, num_val)
				}
			}
		}

		if reader == "soil" {
			CreateMap(line, soil)
		}
		if reader == "fertilizer" {
			CreateMap(line, fertilizer)
		}
		if reader == "water" {
			CreateMap(line, water)
		}
		if reader == "light" {
			CreateMap(line, light)
		}
		if reader == "temperature" {
			CreateMap(line, temperature)
		}
		if reader == "humidity" {
			CreateMap(line, humidity)
		}
		if reader == "location" {
			CreateMap(line, location)
		}
		
	}

	dat.Close()

	// var _soil int
	// var _fert int
	// var _water int
	// var _light int
	// var _temp int
	// var _hum int
	// var _loc int

	// locations := []int{}

	// for _,seed := range seeds {

	// 	if val, ok := soil[seed]; ok {
	// 		_soil = val
	// 	}else {
	// 		_soil = seed
	// 	}

	// 	if val, ok := fertilizer[_soil]; ok {
	// 		_fert = val
	// 	}else {
	// 		_fert = _soil
	// 	}

	// 	if val, ok := water[_fert]; ok {
	// 		_water = val
	// 	}else {
	// 		_water = _fert
	// 	}

	// 	if val, ok := light[_water]; ok {
	// 		_light = val
	// 	}else {
	// 		_light = _water
	// 	}

	// 	if val, ok := temperature[_light]; ok {
	// 		_temp = val
	// 	}else {
	// 		_temp = _light
	// 	}

	// 	if val, ok := humidity[_temp]; ok {
	// 		_hum = val
	// 	}else {
	// 		_hum = _temp
	// 	}

	// 	if val, ok := location[_hum]; ok {
	// 		_loc = val
	// 	}else {
	// 		_loc = _hum
	// 	}

	// 	fmt.Println("Location: ", _loc)
	// 	locations = append(locations, _loc)
	// }
	// min := slices.Min(locations)

	// fmt.Println("Minimum Location: ", min)

	elapsed := time.Since(executionTime)
	fmt.Println("Done in", elapsed)
}

