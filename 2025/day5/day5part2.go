package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	beg  int64
	end  int64
	prev *Node
	next *Node
}

// Use linked list to track where to insert new ranges
// New ranges will replace existing ranges any time the bounds would be extended
func buildRangesCondensed(r string, root *Node) *Node {
	parsed := strings.Split(r, "-")
	beg, err := strconv.ParseInt(parsed[0], 10, 64)
	if err != nil {
		panic(err)
	}
	end, err := strconv.ParseInt(parsed[1], 10, 64)
	if err != nil {
		panic(err)
	}

	if root.beg == 0 && root.end == 0 {
		*root = Node{
			beg:  beg,
			end:  end,
			prev: nil,
			next: nil,
		}
		return root
	}
	cursor := root
	node := &Node{
		beg:  beg,
		end:  end,
		prev: nil,
		next: nil,
	}

	for cursor != nil {
		if node.end < cursor.beg {
			insertBefore(node, cursor)
			break
		}
		if node.beg <= cursor.beg {
			if node.end <= cursor.end {
				node.end = cursor.end
			}
			replace(node, cursor)
			break
		}
		if node.beg >= cursor.beg && node.beg <= cursor.end {
			if node.end <= cursor.end {
				break
			}
			node.beg = cursor.beg
			replace(node, cursor)
		}
		if cursor.next == nil {
			// Insert after, not worth its own function
			node.prev = cursor
			cursor.next = node
			break
		}
		cursor = cursor.next
	}

	if root.prev != nil {
		return root.prev
	}
	return root
}

func insertBefore(node *Node, cursor *Node) {
	if cursor.prev != node {
		node.prev = cursor.prev
	}
	cursor.prev = node
	node.next = cursor
	if node.prev != nil {
		node.prev.next = node
	}

}

func replace(node *Node, cursor *Node) {
	if cursor.prev != node {
		node.prev = cursor.prev
	}
	if cursor.next != node {
		node.next = cursor.next
	}
	if node.prev != nil {
		node.prev.next = node
	}
	if node.next != nil {
		node.next.prev = node
	}
}

// Walk entire linked list, and sum the ranges together
func countRanges(root *Node) int64 {
	var sum int64 = 0
	cursor := root
	for cursor != nil {
		sum += cursor.end - cursor.beg + 1
		cursor = cursor.next
	}
	return sum
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	var root *Node = &Node{}
	for {
		if s, err := r.ReadString('\n'); err == nil {
			if s == "\n" {
				break
			}
			s = s[0 : len(s)-1]
			root = buildRangesCondensed(s, root)
		} else {
			break
		}
	}

	fmt.Println(countRanges(root))
}
