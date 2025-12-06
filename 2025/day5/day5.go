package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func buildRanges(r string, rgs map[int64]int64) {
	parsed := strings.Split(r, "-")
	b1, err := strconv.ParseInt(parsed[0], 10, 64)
	if err != nil {
		panic(err)
	}
	e1, err := strconv.ParseInt(parsed[1], 10, 64)
	if err != nil {
		panic(err)
	}
	if _, exists := rgs[b1]; exists {
		if e1 <= rgs[b1] {
			return
		}
	}
	rgs[b1] = e1
}

func fresh(ing string, rgs map[int64]int64) bool {
	i, err := strconv.ParseInt(ing, 10, 64)
	if err != nil {
		panic(err)
	}
	for b, e := range rgs {
		if i >= b && i <= e {
			return true
		}
	}
	return false
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	sum := 0
	ranges := map[int64]int64{}
	for {
		if s, err := r.ReadString('\n'); err == nil {
			if s == "\n" {
				break
			}
			s = s[0 : len(s)-1]
			buildRanges(s, ranges)
		} else {
			break
		}
	}
	for {
		if s, err := r.ReadString('\n'); err == nil {
			s = s[0 : len(s)-1]
			if fresh(s, ranges) {
				fmt.Printf("%s: is fresh\n", s)
				sum++
			} else {
				fmt.Printf("%s: is spoiled\n", s)
			}
		} else {
			break
		}
	}
	fmt.Println(sum)
}
