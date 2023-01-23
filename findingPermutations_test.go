package main

import "testing"

func BenchmarkFindingPermutations(b *testing.B) {
	x := []int{1, 2, 3, 4, 5}
	answer := [][]int{x}

	for i := 0; i < b.N; i++ {
		findPermutation(x, 0, &answer)
	}
}

func BenchmarkPermutationsStackOverflow(b *testing.B) {
	x := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		permutationsStackOverflow(x)
	}
}
