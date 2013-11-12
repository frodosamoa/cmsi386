package gowarmup

import (
	"testing"
	"sort"
	"strings"
)

// Returns true if two int arrays/slices are equal.

func testIntArrayEq(a, b []int) bool {
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

// Returns true if two string arrays/slices are equal.

func testStringArrayEq(a, b []string) bool {
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

// Returns true if the two passed strings split after UTF-8 sequence into an
// array of those sequences are sorted and equal each other. 

func anagram(a, b string) bool {
	s, t := strings.Split(a, ""), strings.Split(b, "")
	sort.Strings(s)
	sort.Strings(t)
	return testStringArrayEq(s, t)
}


// Returns an int array of a base raised up to succesive powers
// & including the limit. Helper function for testing.

func powersArray(base int, limit int) []int {
	powerArray := []int{}
	powers(base, limit, func(p int) { powerArray = append(powerArray, p) })
	return powerArray
}

// Returns an int array of the powers of two up to & including the limit.
// Helper funciton for testing.

func powersOfTwoArray(limit int) []int {
	powerOfTwoArray := []int{}
	powersOfTwo(limit, func(p int) { powerOfTwoArray = append(powerOfTwoArray, p) })
	return powerOfTwoArray
}

// Helper function for tests that use arrays of structs composed of two int arrays.

func testIntArrayStruct(tests []struct{actual []int; expected []int}, t *testing.T) {
    for i, test := range tests {
		if !testIntArrayEq(test.actual, test.expected) {
			t.Errorf("Test %d: expected %v got %v", i, test.actual, test.expected)
		}
	}
}


// Test are stored in structs for table driven testing. Each one is formatted
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

var scrambleTests = []string{"", "a", "rat", "zzz", "^*&^*&^▱ÄÈËɡɳɷ", "I like turtles", "Happiness depends upon ourselves."}

var powersOfTwoTests = []struct {
		actual 		[]int
		expected 	[]int
}{
	{powersOfTwoArray(-1), []int{}},
    {powersOfTwoArray(0), []int{}},
    {powersOfTwoArray(2), []int{1, 2}},
    {powersOfTwoArray(3), []int{1, 2}},
    {powersOfTwoArray(5), []int{1, 2, 4}},
    {powersOfTwoArray(9), []int{1, 2, 4, 8}},
    {powersOfTwoArray(300), []int{1, 2, 4, 8, 16, 32, 64, 128, 256}},
    {powersOfTwoArray(2500), []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048}},
    {powersOfTwoArray(4194305), []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576, 2097152, 4194304}},
    {powersOfTwoArray(134217729), []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576, 2097152, 4194304, 8388608, 16777216, 33554432, 67108864, 134217728}},
}

var powersTests = []struct {
		actual 		[]int
		expected 	[]int
}{
	{powersArray(0, 0), []int{}},
    {powersArray(0, -1), []int{}},
    {powersArray(-1, -4), []int{}},
    {powersArray(-2, 17), []int{1, -2, 4, -8, 16, -32}},
    {powersArray(3, 400), []int{1, 3, 9, 27, 81, 243}},
    {powersArray(57, 900000000000000), []int{1, 57, 3249, 185193, 10556001, 601692057, 34296447249, 1954897493193, 111429157112001}},
    {powersArray(6, 666), []int{1, 6, 36, 216}},
    {powersArray(-300, 28000000), []int{1, -300, 90000, -27000000}},
}

var interleaveTests = []struct {
		actual 		[]int
		expected 	[]int
}{
	{interleave([]int{}, []int{}), []int{}},
	{interleave([]int{1}, []int{}), []int{1}},
	{interleave([]int{}, []int{1}), []int{1}},
	{interleave([]int{0, 2, 4, 6, 8}, []int{1, 3, 5, 7, 9}), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	{interleave([]int{32}, []int{0, 90, 1, 453, 444}), []int{32, 0, 90, 1, 453, 444}},
	{interleave([]int{0, 90, 1, 453, 444}, []int{32}), []int{0, 32, 90, 1, 453, 444}},
	{interleave([]int{12345, 67890}, []int{98765, 43210}), []int{12345, 98765, 67890, 43210}},
	{interleave([]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{8, 7, 6, 5, 4, 3, 2, 1}), []int{1, 8, 2, 7, 3, 6, 4, 5, 5, 4, 6, 3, 7, 2, 8, 1}},
}

var stutterTests = []struct {
		actual 		[]int
		expected 	[]int
}{
	{stutter([]int{}), []int{}},
	{stutter([]int{0, 1, 2, 3, 4}), []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4}},
	{stutter([]int{123456789}), []int{123456789, 123456789}},
	{stutter([]int{12, 23, 34, 45, 56, 67, 78, 89}), []int{12, 12, 23, 23, 34, 34, 45, 45, 56, 56, 67, 67, 78, 78, 89, 89}},
	{stutter([]int{1342, 1234, 4231}), []int{1342, 1342, 1234, 1234, 4231, 4231}},
	{stutter([]int{1, 2}), []int{1, 1, 2, 2}},
}


// TESTS

func TestChange(t *testing.T) {
	testIntArrayStruct(changeTests, t)
}

func TestStripVowels(t *testing.T) {
	for i, str := range stripVowelsTests {
		if str.actual != str.expected {
			t.Errorf("Test %d: expected %v got %v", i, str.actual, str.expected)
		}
	}
}

func TestScramble(t *testing.T) {
	for i, s := range scrambleTests {
		if !anagram(s, scramble(s)) {
			t.Errorf("Test %d: %s did not scramble correctly", i, s)
		}
	}
}

func TestPowersOfTwo(t *testing.T) {
	testIntArrayStruct(powersOfTwoTests, t)
}

func TestPowers(t *testing.T) {
	testIntArrayStruct(powersTests, t)
}

func TestInterleave(t *testing.T) {
	testIntArrayStruct(interleaveTests, t)
}

func TestStutter(t *testing.T) {
	testIntArrayStruct(stutterTests, t)
}
