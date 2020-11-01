package main 

import (
	"testing"
)

var testInput = [][]int{
	{7, 2, 23, 3, 34, 4 ,54, 43, 69, 13, 34, 8, 19, 12, 10},
	{1, 2, 3, 4, 5, 6, 7},
	{10, 9, 2, 5, 3, 7, 101, 18},
	{5, 89, 10, 4, 13, 11, 23, 26, 45},
	{1, 12, 13, 5, 100, 34, 12, 6},
}

var expectedResult = []int {5, 7, 4, 6, 4}

func TestFindMaxSuccession(t *testing.T) {
	for i, inputArray := range testInput {
		k, _ := FindMaxSuccession(inputArray)
		e := expectedResult[i]
		if k != e {
			t.Errorf("\nError in testInput[%d]:\nExpected: %d\nGot: %d", i, k, e)
		}
	}
}
