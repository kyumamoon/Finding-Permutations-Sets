package main

import "fmt"

func findPermutation(start []int, depth int, permutations *([][]int)) {

	permutationsStorage := [][]int{}

	if depth == len(start) {
		return
	} else {

		for i := depth; i < len(start); i++ {
			permutation := make([]int, len(start))
			copy(permutation, start)
			permutation[depth], permutation[i] = permutation[i], permutation[depth]

			permutationsStorage = append(permutationsStorage, permutation)
			if i != depth {
				*permutations = append(*permutations, permutation)
			}
			//findPermutation(permutation, (depth + 1), permutations)
		}

		for _, e := range permutationsStorage {
			findPermutation(e, (depth + 1), permutations)
		}
	}
}

func findPermutation2(start []int, depth int, permutations *([][]int)) {
	if depth == len(start)-1 {
		return
	} else {
		for i := depth; i < len(start); i++ {
			permutation := start
			permutation[depth], permutation[i] = permutation[i], permutation[depth]
			if i != depth {
				*permutations = append(*permutations, permutation)
			}
			findPermutation(permutation, (depth + 1), permutations)
		}
	}
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	answer := [][]int{x}
	findPermutation2(x, 0, &answer)

	fmt.Println("PERMUTATIONS:", len(answer))

	for _, e := range answer {
		fmt.Println(e)
	}
}
