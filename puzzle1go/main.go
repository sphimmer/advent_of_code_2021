package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Printf("error opening file %v", err)
	}
	inputString := string(bs)

	inputValues := strings.Split(inputString, "\n")
	intValues := convertStringSliceToIntSlice(inputValues)
	fmt.Printf("1a number increased: %d\n", solve1a(intValues))
	fmt.Printf("1b number increased: %d\n", solve1b(intValues))
}

func convertStringSliceToIntSlice(input []string) []int {
	output := []int{}

	for _, v := range input {
		intV, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		output = append(output, intV)
	}
	return output
}

func solve1a(inputValues []int) int {

	count := 0
	for i, v := range inputValues {
		if i == 0 {
			continue
		}

		if v > inputValues[i-1] {
			count++
		}
	}
	return count
}

func solve1b(input []int) int {
	count := 0
	prevSum := input[0] + input[1] + input[2]
	for i := 1; i < len(input); i++ {
		if i+2 == len(input) {
			break
		}
		sum := input[i] + input[i+1] + input[i+2]
		if sum > prevSum {
			count++
		}
		prevSum = sum
	}
	return count
}
