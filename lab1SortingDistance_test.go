package main

import "testing"

func TestCorrectSortingCase(t *testing.T) {
	actual := []string{"abbb", "b", "bb","a"}
	expected := []string{"a", "b", "bb","abbb"}
	keyword := "a"

	result := SortArrayByLevenshteinDistance(actual, keyword)

	if !Equal (expected, result) {
		t.Error("Array distance sorting fail. result: ", result, "; expected: ", expected)
	}
}

func TestIncorrectSortingCase(t *testing.T) {
	actual := []string{"abbb", "b", "bb","a"}
	expected := []string{"b", "a", "bb","abbb"}
	keyword := "a"

	result := SortArrayByLevenshteinDistance(actual, keyword)

	if Equal (expected, result) {
		t.Error("Array distance incorrect case sorting fail. result: ", result, "; expected: ", expected)
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}