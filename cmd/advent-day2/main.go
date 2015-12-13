package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calculateWrappingPaper(l, w, h int) (int, int) {
	var ans [3]int
	inp := [3]int{l, w, h}
	pos := 0
	for i := 0; i < 2; i++ {
		for j := i + 1; j < 3; j++ {
			ans[pos] = inp[i] * inp[j]
			pos++
		}
	}
	sort.Ints(inp[:])

	paper := 0
	min := 0
	for _, n := range ans {
		if min == 0 || n < min {
			min = n
		}
		paper += 2 * n
	}
	paper += min

	ribbon := 2*inp[0] + 2*inp[1] + inp[0]*inp[1]*inp[2]
	return paper, ribbon
}

func parseLine(line string) (l, w, h int, err error) {
	parts := strings.Split(line, "x")
	if len(parts) != 3 {
		err = fmt.Errorf("expected exactly 3 dimensions, got %d", len(parts))
		return
	}

	l, err = strconv.Atoi(parts[0])
	if err != nil {
		err = fmt.Errorf("unable to parse length: %s", err)
		return
	}

	w, err = strconv.Atoi(parts[1])
	if err != nil {
		err = fmt.Errorf("unable to parse width: %s", err)
		return
	}

	h, err = strconv.Atoi(parts[2])
	if err != nil {
		err = fmt.Errorf("unable to parse height: %s", err)
		return
	}
	return
}

func realMain() int {
	var sumPaper, sumRibbon int
	line := 1

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l, w, h, err := parseLine(scanner.Text())
		if err != nil {
			fmt.Printf("line %d: %s: %s\n", line, scanner.Text(), err)
			return 1
		}

		paper, ribbon := calculateWrappingPaper(l, w, h)
		sumPaper += paper
		sumRibbon += ribbon
		line++
	}
	fmt.Println("Sum of Paper:", sumPaper)
	fmt.Println("Sum of Ribbon:", sumRibbon)
	return 0
}

func main() {
	os.Exit(realMain())
}
