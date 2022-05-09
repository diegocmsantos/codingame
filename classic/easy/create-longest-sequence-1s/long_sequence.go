package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	b := scanner.Text()
	fmt.Println(findLongest1s(b))
}

func findLongest1s(b string) int {
	idxs := make([]int, 0)
	idxs = append(idxs, 0)
	for i, s := range b {
		if s == '0' {
			idxs = append(idxs, i)
		}
	}
	idxs = append(idxs, len(b))

	longest := 0
	for i := 0; i < len(idxs)-2; i++ {
		s := b[idxs[i]:idxs[i+2]]
		size := len(s)
		if strings.Count(s, "0") == 2 {
			size -= 1
		}
		if longest < size {
			longest = size
		}
	}
	return longest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
