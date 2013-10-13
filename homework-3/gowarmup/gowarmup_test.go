package gowarmup

import (
	"testing"
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
        {change(8), []int{0, 0, 1, 3}},
        {change(1412873456), []int{56514938, 0, 1, 1}},
        {change(41), []int{1, 1, 1, 1}},
        {change(890), []int{35, 1, 1, 0}},
        {change(4), []int{0, 0, 0, 4}},
        {change(123456789), []int{4938271, 1, 0, 4}},
        {change(250), []int{10, 0, 0, 0}},
        {change(519), []int{20, 1, 1, 4}},
        {change(1111), []int{44, 1, 0, 1}},

        // TODO: Negative change
        // {change(-2), []int{0, 0, 0, 0}},
        // {change(-30000), []int{0, 0, 0, 0}},
}


func TestChange(t *testing.T) {
    for i, change := range changeTests {
		if (!testEq(change.in, change.out)) {
			t.Errorf("Test %d: expected %v got %v", i, change.in, change.out)
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
