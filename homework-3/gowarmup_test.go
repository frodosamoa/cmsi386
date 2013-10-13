package gowarmup

import (
	"testing" 
	"fmt"
)

func testEq(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }

    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}

func testChange(t *testing.T) {
	cents := change(97)
	fmt.Println(cents)
	change := []int{3, 2, 0, 2}
	if (testEq(cents, change)) {
		t.Errorf("Expected [3, 2, 0, 2] got ", change)
	}
}

func testStripVowels(t *testing.T) {
	
}

func testScramble(t *testing.T) {
	
}

func testPowersOfTwo(t *testing.T) {
	
}

func testPowers(t *testing.T) {
	
}

func testInterleave(t *testing.T) {
	
}

func testStutter(t *testing.T) {
	
}
