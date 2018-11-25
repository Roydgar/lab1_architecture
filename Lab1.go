package main

import (
	"sort"
)

func LevenshteinDistance(s1, s2 string) int {
	min := func(values ...int) int {
		m := values[0]
		for _, v := range values {
			if v < m {
				m = v
			}
		}
		return m
	}
	r1, r2 := []rune(s1), []rune(s2)
	n, m := len(r1), len(r2)
	if n > m {
		r1, r2 = r2, r1
		n, m = m, n
	}
	currentRow := make([]int, n+1)
	previousRow := make([]int, n+1)
	for i := range currentRow {
		currentRow[i] = i
	}
	for i := 1; i <= m; i++ {
		for j := range currentRow {
			previousRow[j] = currentRow[j]
			if j == 0 {
				currentRow[j] = i
				continue
			} else {
				currentRow[j] = 0
			}
			add, del, change := previousRow[j]+1, currentRow[j-1]+1, previousRow[j-1]
			if r1[j-1] != r2[i-1] {
				change++
			}
			currentRow[j] = min(add, del, change)
		}
	}
	return currentRow[n]
}

func SortArrayByLevenshteinDistance(array []string, keyword string) []string{

	sort.Slice(array, func(i, j int) bool {
		return LevenshteinDistance(array[i], keyword) <  LevenshteinDistance(array[j], keyword)
	})

	return array
}

