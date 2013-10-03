package main

import (
	"fmt"
	"regexp"
	"math/rand"
	"strings"
	"time"
	//"builtin"
)

// Returns a tuple containing, respectively, the smallest number of U.S. quarters, dimes, nickels, and pennies that equal the given amount.

func change(cents int) []int {
	change, denoms, remaining := make([]int, 4), [4]int{25, 10, 5, 1}, cents

	change[0] = remaining / denoms[0]
	remaining %= denoms[0]
	change[1] = remaining / denoms[1]
	remaining %= denoms[1]
	change[2] = remaining / denoms[2]
	change[3] = remaining % denoms[2]

	return change
}

// Returns the string which is passed but with all vowels removed.

func stripVowels(s string) string {
	return regexp.MustCompile(`(?i)[aeiou]`).ReplaceAllString(s, "")
}

// Randomly permutes a string. 

func scramble(s string) string {
    r, c := 0, ""
    a := strings.Split(s, "")

    for i := len(a); i > 0; i -= 1 {
        r = random(0, i)
        c = a[i-1];
        a[i-1] = a[r];
        a[r] = c;
    }

    return strings.Join(a, "");
}

// Random int generator between min and max. 

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

// Yields successive powers of two starting at 1 and going up to some limit, consumed with a callback.

func powersOfTwo(limit int, f func(int) int) {
    for i := 1; i <= limit; i *= 2 {
        f(i)
    }	
}

// Yields powers of an arbitrary base starting at exponent 0 and going up to some limit, consumed with a callback.

func powers(base int, limit int, f func(int) int) {
    if base == 1 {
        f(base)
    } else {
        for i := 1; i <= limit; i *= base {
            f(i)
        }
    }	
}

// Interleaves two slices.

func interleave(a []int, b []int) []int {
	c := []int{}

	for i := 0; i < len(c); i++ {
		if i < len(a) {
			c = append(c, 3)
		}

		if i < len(b) {
			c = append(c, b[i])
		}
	}

	return c
}

// Stutters a sequence, i.e. (4, 3) ⇒ (4, 4, 3, 3)

func stutter(a []int) []int {
	return interleave(a, a)
}

func print(x int) int {
	return x
}

func main() {
	fmt.Println("poop")
	a := []int{1, 2, 3}
	b := []int{4, 5, 6, 7, 8}
	c := interleave(a, b)
	fmt.Println(c)
	powersOfTwo(800, print)
}
