package main

import "fmt"

func findPermutation(start []int, depth int, permutations *([][]int)) {

	if depth == len(start) {
		return
	} else {
		permutationsStorage := [][]int{}
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

func findPermutationInterface(start []interface{}, depth int, permutations *([][]interface{})) {

	if depth == len(start) {
		return
	} else {
		permutationsStorage := [][]interface{}{}
		for i := depth; i < len(start); i++ {
			permutation := make([]interface{}, len(start))
			copy(permutation, start)
			permutation[depth], permutation[i] = permutation[i], permutation[depth]

			permutationsStorage = append(permutationsStorage, permutation)
			if i != depth {
				*permutations = append(*permutations, permutation)
			}
			//findPermutation(permutation, (depth + 1), permutations)
		}

		for _, e := range permutationsStorage {
			findPermutationInterface(e, (depth + 1), permutations)
		}
	}
}

func findPermutation2(start []int, depth int, permutations *([][]int)) {
	if depth == len(start) {
		return
	} else {
		for i := depth; i < len(start); i++ {
			//permutation := start
			permutation := make([]int, len(start))
			copy(permutation, start)
			permutation[depth], permutation[i] = permutation[i], permutation[depth]
			if i != depth {
				*permutations = append(*permutations, permutation)
			}
			findPermutation2(permutation, (depth + 1), permutations)
		}
	}
}

func main() {
	//x := []interface{}{1, "a", true}
	x := []int{1}
	answer := [][]int{x}
	findPermutation2(x, 0, &answer)

	for _, e := range answer {
		fmt.Println(e)
	}

	fmt.Println("PERMUTATIONS:", len(answer))
	//fmt.Println(answer)
}
