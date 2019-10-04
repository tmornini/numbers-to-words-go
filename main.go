package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())

		if err != nil {
			panic(err)
		}

		fmt.Println(
			integerToWords(value),
		)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}
