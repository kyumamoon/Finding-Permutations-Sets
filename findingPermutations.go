package main

import "fmt"

func findPermutation(start []int, depth int, permutations *([][]int)) {

	permutationsStorage := [][]int{}

	if depth == len(start) {
		return
	} else {

		fmt.Println(start, "DEPTH:", depth)
		if depth == 0 {
			*permutations = append(*permutations, start)

			findPermutation(start, (depth + 1), permutations)
		}

		for i := depth + 1; i < len(start); i++ {
			permutation := make([]int, len(start))
			copy(permutation, start)
			permutation[depth], permutation[i] = permutation[i], permutation[depth]

			permutationsStorage = append(permutationsStorage, permutation)
			*permutations = append(*permutations, permutation)
		}

		for _, e := range permutationsStorage {
			findPermutation(e, (depth + 1), permutations)
		}
	}
}

func main() {
	x := []int{1, 2, 3}
	answer := [][]int{}
	findPermutation(x, 0, &answer)

	fmt.Println("PERMUTATIONS:", len(answer))
	fmt.Println(answer)
	//fmt.Println(answer)
}
