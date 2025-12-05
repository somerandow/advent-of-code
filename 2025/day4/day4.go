package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	rowLen := 0
	count := 0
	rolls := []bool{}
	for {

		if s, err := r.ReadByte(); err == nil {
			switch s {
			case '\n':
				if rowLen == 0 {
					rowLen = count
				}
			case '@':
				rolls = append(rolls, true)
			case '.':
				rolls = append(rolls, false)
			}
			count++
		} else {
			break
		}
	}
	fmt.Println(accessiblePart2(rolls, rowLen))
}
func accessiblePart2(rolls []bool, rowLen int) int {
	total := 0
	for {
		sum := accessible(rolls, rowLen, true)
		if sum == 0 {
			break
		}
		total += sum
	}
	return total
}

func accessible(rolls []bool, rowLen int, remove bool) int {
	total := 0
	graph := []string{}
	for i := 0; i < len(rolls); i++ {
		if !rolls[i] {
			graph = append(graph, ".")
			continue
		}
		sum := 0
		for j := -1; j < 2; j++ {
			for k := -1; k < 2; k++ {
				if j == 0 && k == 0 {
					continue
				}
				idx := i + rowLen*j + k
				// OOB check
				if idx < 0 || idx > len(rolls)-1 {
					continue
				}
				if i%rowLen == 0 && k == -1 {
					continue
				}
				// Wrap around check, left side
				if i%rowLen == 0 && k == -1 {
					continue
				}
				// Wrap around check, right side
				if i%rowLen == rowLen-1 && k == 1 {
					continue
				}
				if rolls[idx] {
					sum++
				}
			}
		}
		if sum < 4 {
			total++
			graph = append(graph, "X")
			if remove {
				rolls[i] = false
			}
		} else {
			graph = append(graph, "@")
		}
	}
	for i := 0; i <= len(graph)-rowLen; i += rowLen {
		fmt.Println(strings.Join(graph[i:i+rowLen], ""))
	}
	return total
}
