package main

import (
	"fmt"
	"io"
	"os"
)

func realMain() int {
	count := 0
	pos := 0
	basement := 0

	data := make([]byte, 4096)
	for {
		n, err := os.Stdin.Read(data)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
				return 1
			}
			break
		}

		for i := 0; i < n; i++ {
			switch data[i] {
			case '(':
				count++
			case ')':
				count--
			}

			pos++
			if basement == 0 && count < 0 {
				basement = pos
			}
		}
	}

	fmt.Println("Final Floor:", count)
	if basement > 0 {
		fmt.Println("First Basement Floor:", basement)
	}
	return 0
}

func main() {
	os.Exit(realMain())
}
