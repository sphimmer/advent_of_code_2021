package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := parseFile()
	inputValues := strings.Split(input, "\n")

	fmt.Println("3A: ", solve3a(inputValues))
	fmt.Println("3B: ", solve3b(inputValues))
}

func parseFile() string {
	bs, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Printf("error opening file %v", err)
	}
	inputString := string(bs)
	return inputString
}

func solve3a(input []string) int64 {

	bitCounters := []bitCounter{}
	for i := 0; i < len(input[0]); i++ {
		bitCounters = append(bitCounters, bitCounter{one: 0, zero: 0})
	}

	for _, v := range input {
		for i := range v {
			value, err := strconv.Atoi(string((v[i])))
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
			if value == 0 {
				bitCounters[i].zero++
			} else if value == 1 {
				bitCounters[i].one++
			}
		}
	}
	// fmt.Println(bitCounters)
	gammaRate := []string{}
	epsilonRate := []string{}
	for _, bc := range bitCounters {
		gammaRate = append(gammaRate, bc.highValue())
		epsilonRate = append(epsilonRate, bc.lowValue())
	}
	binaryStringGR := strings.Join(gammaRate, "")
	binaryStringER := strings.Join(epsilonRate, "")
	// fmt.Println(binaryStringGR, binaryStringER)
	grInt, err := strconv.ParseInt(binaryStringGR, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	erInt, err := strconv.ParseInt(binaryStringER, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(grInt, erInt)
	return grInt * erInt
}

func solve3b(input []string) int64 {
	binaryStringOR := filter(input, 0, true)
	binaryStringCSR := filter(input, 0, false)
	oxygenGenRating, err := strconv.ParseInt(binaryStringOR, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	co2ScrubRating, err := strconv.ParseInt(binaryStringCSR, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return oxygenGenRating * co2ScrubRating
}

func filter(input []string, whichBit int, countMost bool) string {
	haveOne := []string{}
	haveZero := []string{}
	for _, v := range input {
		value := string((v[whichBit]))
		if value == "1" {
			haveOne = append(haveOne, v)
		} else if value == "0" {
			haveZero = append(haveZero, v)
		}
	}

	whichBit++
	var keep []string

	if countMost {
		if len(haveOne) >= len(haveZero) {
			keep = haveOne
		} else {
			keep = haveZero
		}
	} else {
		if len(haveOne) >= len(haveZero) {
			keep = haveZero
		} else {
			keep = haveOne
		}
	}
	if len(keep) == 1 {
		return keep[0]
	} else if whichBit < len(input[0]) {
		return filter(keep, whichBit, countMost)
	} else {
		return keep[0]
	}
}
