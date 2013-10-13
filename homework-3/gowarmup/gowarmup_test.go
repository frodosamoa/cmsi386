package gowarmup

import (
	"testing"
	"fmt"
)

// Returns true if two int arrays/slices are equal.

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

// Test are stored in struct sfor table driven testing. Each one is formatted
// as so. Actual output, then the expected output, denoted by those respecive
// words.

var changeTests = []struct {
        actual  	[]int
        expected 	[]int
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

var stripVowelsTests = []struct {
		actual 		string
		expected 	string
}{
	{stripVowels("Hello, world"), "Hll, wrld"},
	{stripVowels("I like turtles"), " lk trtls"},
	{stripVowels("aeiou"), ""},
	{stripVowels("bcdfghjklmnpqrstvwxyz"), "bcdfghjklmnpqrstvwxyz"},
	{stripVowels("abecidoeuf"), "bcdf"},
	{stripVowels("Happiness depends upon ourselves"), "Hppnss dpnds pn rslvs"},
	{stripVowels("Chrononhotonthologos"), "Chrnnhtnthlgs"},
	{stripVowels("HONORIFICABILITUDINITATIBUS"), "HNRFCBLTDNTTBS"},
	{stripVowels("Philip M. Dorin, John David N. Dionisio, Ray Toal, Masao Kitamura"), "Phlp M. Drn, Jhn Dvd N. Dns, Ry Tl, Ms Ktmr"},
	{stripVowels("`~!@#$^&*()_+-=[]{};./>?"), "`~!@#$^&*()_+-=[]{};./>?"},
}

var interleaveTests = []struct {
		actual 		[]int
		expected 	[]int
}{
	{interleave([]int{},[]int{}), []int{}},
	{interleave([]int{0, 2, 4, 6, 8}, []int{1, 3, 5, 7, 9}), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	// {interleave(), },
	// {interleave(), },
	// {interleave(), },
	// {interleave(), },
	// {interleave(), },
	// {interleave(), },
	// {interleave(), },
	// {interleave(), },
}

var stutterTests = []struct {
		actual 		[]int
		expected 	[]int
}{
	// {stutter(), },
	// {stutter(), },
	// {stutter(), },
	// {stutter(), },
	// {stutter(), },
	// {stutter(), },
	// {stutter(), },
	// {stutter(), },
	// {stutter(), },
	// {stutter(), },
}

func testIntArrayStructs(tests []struct{}, t *testing.T) {
	fmt.Println(tests)
    for i, test := range tests {
		if (!testEq(test.actual, test.expected)) {
			t.Errorf("Test %d: expected %v got %v", i, test.actual, test.expected)
		}
	}
}

// TESTS

func TestChange(changeTests []struct{}, t *testing.T) {
	testIntArrayStructs(changeTests, t)
}

func TestStripVowels(t *testing.T) {
	for i, str := range stripVowelsTests {
		if (str.actual != str.expected) {
			t.Errorf("Test %d: expected %v got %v", i, str.actual, str.expected)
		}
	}
}

func TestScramble(t *testing.T) {
	
}

func TestPowersOfTwo(t *testing.T) {
	
}

func TestPowers(t *testing.T) {
	
}

func TestInterleave(interleaveTests []struct{}, t *testing.T) {
	testIntArrayStructs(interleaveTests, t)
}

func TestStutter(stutterTests []struct{}, t *testing.T) {
	testIntArrayStructs(stutterTests, t)
}
