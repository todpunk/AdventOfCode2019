package main

import (
	"bufio"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"math"
	"os"
	"strconv"
)

var (
	day   = kingpin.Arg("day", "Advent day to run").Int()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	fmt.Printf("Would run day: %d \n", *day)
	switch *day {
	case 1:
		day1()
	default:
		fmt.Printf("We don't have that day...")
	}
}

func day1() {
	file, err := os.Open("./day1-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var modules []int64
	var sum int64 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var module int64
		module, _ = strconv.ParseInt(scanner.Text(), 10, 32)
		sum += int64(math.Floor(float64(module / 3.0))) - 2
		modules = append(modules, module)
	}

	fmt.Printf("Total fuel required: %d \n", sum)
	fmt.Printf("But we need to fuel that fuel, so....\n")

	var newFuel int64
	sum = 0
	for i, _ := range modules {
		for newFuel = modules[i]; newFuel > 0; {
			newFuel = int64(math.Floor(float64(newFuel / 3.0))) - 2
			if newFuel > 0 {
				sum += newFuel
			}
		}
	}
	fmt.Printf("New Total fuel required (calculated per module): %d \n", sum)
}