package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxJoltagePart2(s string) string {
	// Trim newline
	s = s[:len(s)-1]
	digits := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i, c := range s {
		next := int(c - 0x30)
		for k := 0; k < len(digits); k++ {
			if next > digits[k] && i < len(s)-11+k {
				digits[k] = next
				// Clear all remaining digits
				for j := k + 1; j < len(digits); j++ {
					digits[j] = 0
				}
				break
			}
		}
	}
	output := fmt.Sprintf("%s", strings.Join(func(i []int) []string {
		output := []string{}
		for _, c := range i {
			output = append(output, strconv.Itoa(c))
		}
		return output
	}(digits), ""))
	fmt.Printf("Bank: %s\nMax Joltage:%s\n", s, output)
	return output
}

func maxJoltage(s string) string {
	// Trim newline
	s = s[:len(s)-1]
	a := 0
	b := 0
	for i, c := range s {
		next := int(c - 0x30)
		if next > a && i < len(s)-1 {
			a = next
			b = 0
		} else {
			if next > b {
				b = next
			}
		}
	}
	output := fmt.Sprintf("%d%d", a, b)
	fmt.Printf("Bank: %s\nMax Joltage:%s\n", s, output)
	return output
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	sum := 0
	for {
		if s, err := r.ReadString('\n'); err == nil {
			res, _ := strconv.Atoi(maxJoltagePart2(s))
			sum += res
		} else {
			break
		}
	}
	fmt.Println(sum)
}
