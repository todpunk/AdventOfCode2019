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
	fmt.Println("Would run day: %d", *day)
	switch *day {
	case 1:
		day1()
	case 2:
		day2()
	default:
		fmt.Println("We don't have that day...")
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

	fmt.Printf("Total fuel required: %d\n", sum)
	fmt.Println("But we need to fuel that fuel, so....")

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
	fmt.Printf("New Total fuel required (calculated per module): %d\n", sum)
}

func runDay2Comp(codes []int64 ) (pos0 int64 ) {
	for i := 0; i < len(codes)-1; i += 4 {
		opcode := codes[i]
		switch opcode{
		case 99:
			i = len(codes)
		case 1:
			a := codes[codes[i+1]]
			b := codes[codes[i+2]]
			codes[codes[i+3]] = a + b
		case 2:
			a := codes[codes[i+1]]
			b := codes[codes[i+2]]
			codes[codes[i+3]] = a * b
		default:
			i = len(codes)
			fmt.Println("This went poorly")
		}
	}
	return codes[0]
}

func day2() {
	rawcodes := []int64{1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,9,1,19,1,19,5,23,1,23,6,27,2,9,27,31,1,5,31,35,1,35,10,39,1,39,10,43,2,43,9,47,1,6,47,51,2,51,6,55,1,5,55,59,2,59,10,63,1,9,63,67,1,9,67,71,2,71,6,75,1,5,75,79,1,5,79,83,1,9,83,87,2,87,10,91,2,10,91,95,1,95,9,99,2,99,9,103,2,10,103,107,2,9,107,111,1,111,5,115,1,115,2,119,1,119,6,0,99,2,0,14,0}
	var codes = make([]int64, len(rawcodes))
	copy(codes, rawcodes)
	fmt.Println(rawcodes)
	fmt.Println(codes)
	codes[1] = 12
	codes[2] = 2
	fmt.Printf("Position 0: %d\n", runDay2Comp(codes))

	var pos0, noun, verb int64
	for noun = 1; noun < 100; noun++ {
		for verb = 1; verb < 100; verb++ {
			copy(codes, rawcodes)
			codes[1] = noun
			codes[2] = verb
			pos0 = runDay2Comp(codes)
			if pos0 == 19690720 {
				fmt.Printf("Noun: %d Verb: %d Result: %d\n", noun, verb, (100 * noun + verb))
				noun = 99
				verb = 99
			}
		}
	}
}