package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"
)

func numberOfLines() int {
	scanner, count := bufio.NewScanner(os.Stdin), 0

	for scanner.Scan() { 
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 && line[0] != '#' {
			count += 1
		} 
	}

	if err := scanner.Err(); err != nil {
	    log.Fatal(err)
	}

	return count
}

/* Reports the number of non-blank, non-commented lines in the
   file named by the first argument. */

func main() {
	fmt.Println(numberOfLines())
}