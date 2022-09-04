package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string = os.Args[1]

	var cnt int = strings.Count(input, " ")

	if input != "" {
		cnt++
	}

	fmt.Print(cnt)
}
