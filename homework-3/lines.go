package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
)

// Figure out way to pass the first commandline argument in to the top function

func numberOfLines() int {
	scanner, count := bufio.NewScanner(os.Stdin), 0

	for scanner.Scan() { 
		if len(scanner.Text()) > 0 { 
			if scanner.Text()[0] != '#' {
				count += 1
			}
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