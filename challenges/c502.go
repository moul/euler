package challenges

import (
	"fmt"
	"strings"
)

func init() {
	Challenges[502] = Challenge{
		Title:    "Counting Castles",
		Correct:  false,
		Callback: C502,
		ID:       502,
	}
}

type c502Permutation []int

func (p *c502Permutation) Blocks() int {
	ret := 0
	for _, size := range *p {
		if size > 0 {
			ret++
		}
	}
	return ret
}

type c502Permutations map[string]c502Permutation

func (p *c502Permutation) Hash() string {
	ret := ""
	for _, size := range *p {
		ret += fmt.Sprintf("%d", size)
	}
	return ret
}

var cache502 struct {
	Sum          map[int]map[int]map[bool]int
	Permutations map[int]c502Permutations
}

func init() {
	cache502.Sum = make(map[int]map[int]map[bool]int, 0)
	cache502.Permutations = make(map[int]c502Permutations)
}

func c502ComputePermutations(w int) c502Permutations {
	// this can be done more easily by walking from 0 to w^2 in binary
	if w == 0 {
		return c502Permutations{}
	}

	if val, found := cache502.Permutations[w]; found {
		return val
	}

	permutations := c502Permutations{}

	for start := 0; start <= w; start++ {
		startPermutation := c502Permutation{}
		for i := 0; i < start; i++ {
			startPermutation = append(startPermutation, 0)
		}
		if start == w {
			permutations[startPermutation.Hash()] = startPermutation
			break
		}

		for size := 1; size <= w-start; size++ {
			sizePermutation := make(c502Permutation, len(startPermutation))
			copy(sizePermutation, startPermutation)
			sizePermutation = append(sizePermutation, size)

			if start+size == w {
				permutations[sizePermutation.Hash()] = sizePermutation
				continue
			}

			sizePermutation = append(sizePermutation, 0)
			if start+size+1 == w {
				permutations[sizePermutation.Hash()] = sizePermutation
				continue
			}

			// fmt.Printf("w=%d start=%d size=%d sub=%d permutation=%v\n", w, start, size, w-start-size-1, sizePermutation)

			for _, subPermutation := range c502ComputePermutations(w - start - size - 1) {
				concatPermutation := make(c502Permutation, len(sizePermutation))
				copy(concatPermutation, sizePermutation)
				concatPermutation = append(concatPermutation, subPermutation...)
				permutations[concatPermutation.Hash()] = concatPermutation
			}
		}
	}

	cache502.Permutations[w] = permutations
	return permutations
}

func c502ComputeSum(w, h int) int {
	return c502ComputeSumWithBlocks(w, h, 1)
}

func c502ComputeSumWithBlocks(w, h, blocks int) int {
	// one line height only have one possibility
	if h == 1 {
		if blocks%2 == 0 {
			return 1
		}
		return 0
	}

	if val, found := cache502.Sum[w][h][blocks%2 == 0]; found {
		fmt.Printf("%3d.%-3d => %4d  -  %5v  -  (cache)\n", w, h, val, (blocks%2 == 0))
		return val
	}

	ret := 0

	permutations := c502ComputePermutations(w)
	if h == 2 {
		for _, permutation := range permutations {
			even := (blocks+permutation.Blocks())%2 == 0
			// fmt.Println(w, h, even, permutation)
			if even {
				ret++
			}
		}
	} else {
		for _, permutation := range permutations {
			for _, size := range permutation {
				if size > 0 {
					ret += c502ComputeSumWithBlocks(size, h-1, blocks+permutation.Blocks())
				}
			}
		}
	}

	fmt.Printf("%3d.%-3d => %4d  -  %5v  -  (new)\n", w, h, ret, (blocks%2 == 0))
	if _, found := cache502.Sum[w]; !found {
		cache502.Sum[w] = make(map[int]map[bool]int, 0)
	}
	if _, found := cache502.Sum[w][h]; !found {
		cache502.Sum[w][h] = make(map[bool]int, 0)
	}
	cache502.Sum[w][h][blocks%2 == 0] = ret
	return ret
}

func C502() (interface{}, error) {
	/*
		fmt.Println(c502ComputeSum(4, 2))
		fmt.Println(strings.Repeat("-", 70))
		fmt.Println(c502ComputeSum(13, 10))
		fmt.Println(strings.Repeat("-", 70))
		fmt.Println(c502ComputeSum(10, 13))
		fmt.Println(strings.Repeat("-", 70))
	*/
	fmt.Println(c502ComputeSum(3, 3))
	fmt.Println(strings.Repeat("-", 70))

	if true {
		fmt.Println()
		fmt.Println("Sum")
		fmt.Println("===")
		for w, wItems := range cache502.Sum {
			for h, hItems := range wItems {
				for parity, items := range hItems {
					fmt.Println(w, h, parity, items)
				}
			}
		}

		fmt.Println()
		fmt.Println("Permutations")
		fmt.Println("============")
		for i := 0; i <= len(cache502.Permutations); i++ {
			permutations := c502ComputePermutations(i)
			fmt.Printf("permutations(%d) len=%d permutations=%v\n", i, len(permutations), permutations)
		}
	}
	return nil, nil
}
