package main

import (
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		println(line)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}
