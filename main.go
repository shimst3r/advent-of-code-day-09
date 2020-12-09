package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	maxUint = ^uint(0)
	maxInt  = int(maxUint >> 1)
	minInt  = -maxInt - 1
)

func main() {
	inputPath := flag.String("input-path", "test_input.txt", "Path to input file")
	preambleLength := flag.Int("preamble-length", 5, "Length of cypher preamble")
	flag.Parse()
	file, err := os.Open(*inputPath)
	if err != nil {
		panic(err)
	}
	input, err := parseInput(file)
	if err != nil {
		panic(err)
	}
	for idx := 0; idx < len(input)-*preambleLength; idx++ {
		preamble := input[idx : idx+*preambleLength]
		number := input[idx+*preambleLength]
		if !numberIsValid(preamble, number) {
			log.Printf("number %d is not valid\n", number)
			min, max := findEncryptionWeakness(input, number)
			log.Printf("the encryption weakness is: %d + %d = %d\n", min, max, min+max)
			return
		}
	}
}

func parseInput(r io.Reader) ([]int64, error) {
	scanner := bufio.NewScanner(r)
	var result []int64
	for scanner.Scan() {
		number, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)
		if err != nil {
			return []int64{}, err
		}
		result = append(result, number)
	}
	return result, nil
}

func numberIsValid(preamble []int64, number int64) bool {
	for idx, xs := range preamble {
		for idy, ys := range preamble {
			if idx != idy && xs+ys == number {
				return true
			}
		}
	}
	return false
}

func findEncryptionWeakness(input []int64, number int64) (int64, int64) {
	var idx, idy int
outerloop:
	for idx = range input {
	innerloop:
		for idy = idx + 1; idy < len(input); idy++ {
			if sum(input[idx:idy+1]) > number {
				break innerloop
			}
			if sum(input[idx:idy+1]) == number {
				break outerloop
			}
		}
	}
	slice := input[idx : idy+1]
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	return slice[0], slice[len(slice)-1]
}

func sum(slice []int64) (result int64) {
	for _, value := range slice {
		result += value
	}
	return result
}
