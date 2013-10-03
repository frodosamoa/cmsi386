package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
)

func numberOfLines() int {
	scanner, count, block := bufio.NewScanner(os.Stdin), 0, false

	for scanner.Scan() {

		// current line non empty?
		if len(scanner.Text()) > 0 {
			// current line longer than one character?
			if len(scanner.Text()) > 1 {
				// does the first cahracter indicate a potential comment?
				if scanner.Text()[0] == '/' {
					if scanner.Text()[1] != '/' { // non empty non commented line starting with '/''
						count += 1
					} else if scanner.Text()[1] == '*' { // block comment
						block = true
					} 
				// or does the first character a '*' indicating the end of a comment block?
				} else if scanner.Text()[0] == '*' {
					if block && scanner.Text()[1] == '/' {
					    block = false
					}
				} else { // non empty non commented line
					count += 1
				}
			}
		} 
	}

	if err := scanner.Err(); err != nil {
	    log.Fatal(err)
	}

	return count
}

func main() {
	fmt.Println(numberOfLines())
}