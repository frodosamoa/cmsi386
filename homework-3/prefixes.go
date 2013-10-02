package main

import (
	"fmt"
	"os"
	"strings"
)

// Writes successive prefixes of its first input argument, one per line.

func prefixes(s string) {
	pre := strings.Split(s, "")
	for i := 0; i < len(pre) + 1; i++ {
		fmt.Println(strings.Join(pre[:i], ""))
	}
}

func main() {
	prefixes(os.Args[1])
}