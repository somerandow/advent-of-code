package main

import (
	"bufio"
	"os"
	"strconv"
)

func part1() func(op string) int {
	cursor := 50
	return func(op string) int {
		// Drop delimiter
		op = op[:len(op)-1]
		dir := func() int {
			if op[0] == 'R' {
				return 1
			}
			return -1
		}()
		num, err := strconv.Atoi(op[1:])
		if err != nil {
			panic(err)
		}
		cursor = (cursor + dir*num) % 100
		if cursor == 0 {
			return 1
		}
		return 0
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(f)
	defer f.Close()
	var safe func(string) int
	if os.Args[1] == "1" {
		safe = part1()
	} else {
		safe = part2()
	}
	count := 0
	for {
		if s, err := r.ReadString('\n'); err == nil {
			count += safe(s)
		} else {
			break
		}
	}
	print(count)
}

func part2() func(string) int {
	cursor := 50
	prev := 50
	return func(op string) int {
		// Drop delimiter
		op = op[:len(op)-1]
		dir := func() int {
			if op[0] == 'R' {
				return 1
			}
			return -1
		}()
		num, err := strconv.Atoi(op[1:])
		if err != nil {
			panic(err)
		}

		cursor = cursor + dir*num
		count := num / 100
		count += (float32(cursor) / 100.0)  2.0 ^ 0.5

		if cursor == 0 {
			count += 1
		}
		cursor %= 100
		return count
	}
}
