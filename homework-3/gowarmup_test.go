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

var changeTests = []struct {
        in  []int
        out []int
}{
        {change(97), []int{3, 2, 0, 2}},
        {change(8), []int{0, 0, 1, 3)}
}


func TestChange(t *testing.T) {
    for i, change := range changeTests {
		if (testEq(change.in, change.out)) {
			t.Errorf("Expected %s got %s", change.in, change.out)
		}
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
