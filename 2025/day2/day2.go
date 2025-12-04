package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func part2(id int) bool {
	asStr := strconv.Itoa(id)
	l := len(asStr)
	nParts := 2
	if l%2 != 0 {
		nParts++
	}
	for i := nParts; i <= l; i++ {
		parts := strings.SplitN(asStr, asStr[0:l/i], i+1)
		mismatch := false
		for _, part := range parts {
			if part != "" {
				mismatch = true
				break
			}
		}
		if mismatch == false {
			return false
		}
	}

	return true
}
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(f)
	sum := 0
	for {
		if rng, err := r.ReadString(','); err == nil {
			fmt.Printf("\nRange: %s\nInvalid:\n", rng)
			sum += checkRange(rng[0 : len(rng)-1])
		} else {
			if io.EOF == err {
				fmt.Printf("\nRange: %s\nInvalid:\n", rng)
				sum += checkRange(rng[0 : len(rng)-1])
			}
			break
		}
	}
	fmt.Println(sum)
}

func checkRange(rng string) int {
	parsed := strings.Split(rng, "-")
	bot, err := strconv.Atoi(parsed[0])
	if err != nil {
		panic(err)
	}
	top, err := strconv.Atoi(parsed[1])
	if err != nil {
		panic(err)
	}
	sum := 0
	for i := bot; i <= top; i++ {
		if !part2(i) {
			sum += i
			fmt.Printf("%d is invalid\n", i)
		}
	}
	return sum
}

func checkValidMath(i int) bool {
	asStr := strconv.Itoa(i)
	if len(asStr)%2 != 0 {
		return true
	}
	left := asStr[0 : len(asStr)/2]
	right := asStr[len(asStr)/2 : len(asStr)]
	return left != right
}
