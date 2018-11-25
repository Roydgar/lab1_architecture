package main

import (
	"testing"
)
type testpair struct {
	value string
	distance int
	keyword string
}

var kittenKeyword = "kitten"
var singleCharKeyword = "a"

var correctTestCases = []testpair{
	{kittenKeyword, 0, kittenKeyword},
	{"sitting", 3, kittenKeyword},
	{"kitton", 1, kittenKeyword},
	{"kittek", 1, kittenKeyword},
	{"sdf", 6, kittenKeyword},
	{singleCharKeyword, 0, singleCharKeyword},
	{"b", 1, singleCharKeyword},
	{"bb", 2, singleCharKeyword},
	{"bbb", 3, singleCharKeyword},
	{"ab", 1, singleCharKeyword},
	{"aa", 1, singleCharKeyword},
}

var incorrectTestCases = []testpair{
	{"a", 1, singleCharKeyword},
	{"b", 0, singleCharKeyword},
	{"bb", 1, singleCharKeyword},
	{"bbb", 2, singleCharKeyword},
	{"ab", 5, singleCharKeyword},
}

func TestDistanceCorrectCases(t *testing.T) {
	for _, pair := range correctTestCases {
		dist := LevenshteinDistance(pair.value, pair.keyword)
		if dist != pair.distance {
			t.Error("Correct cases test: ",
				"For", pair.value,
				"expected distance ", pair.distance,
				"got", dist,
			)
		}
	}
}

func TestDistanceIncorrectCases(t *testing.T) {
	for _, pair := range incorrectTestCases {
		dist := LevenshteinDistance(pair.value, pair.keyword)
		if dist == pair.distance {
			t.Error("Incorrect cases test: ",
				"For", pair.value,
				"expected distance ", pair.distance,
				"got", dist,
			)
		}
	}
}