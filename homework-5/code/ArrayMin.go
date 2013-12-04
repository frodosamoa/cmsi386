package main

import "fmt"

const MaxInt = int(^uint(0) >> 1)

func ArrayMin(a []int) int {
	return MinFinder(a, MaxInt)
}

func MinFinder(a []int, min int) int {	
	if (len(a) == 0) {
		return min
	} else {
		if (a[0] < min) {
			min = a[0]
		}
	}
	return MinFinder(a[1:], min)
}

func main() {
	a := []int{5, 6, 7, 12, 3, 1}
	b := []int{-432, 3, 1, 234}
	c := []int{-12, -11, -4, 1344234}
	fmt.Println(ArrayMin(a), ArrayMin(b), ArrayMin(c))
}