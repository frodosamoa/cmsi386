package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
)

func numberOfLines() int {
	scanner, count, block := bufio.NewScanner(os.Stdin), 0, false

	for scanner.Scan() { // while we have more lines
		lineLen := len(scanner.Text())
		if lineLen > 1 { // two or more chars
			firstChar, secondChar := scanner.Text()[0], scanner.Text()[1]
			if firstChar == '/' { // potential comment?
				if secondChar == '*' { // block comment
					block = true
				} else if !block && secondChar != '/' { // not comment line
					count += 1
				} 
			} else if firstChar == '*' && secondChar == '/' { // end of comment block?
				block = false
			} else {	
				count += 1
			}
		} else if !block && lineLen > 0 {
			count += 1
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