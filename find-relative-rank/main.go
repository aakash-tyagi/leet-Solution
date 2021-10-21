package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {

	res := make([]string, len(score))
	mapRes := make(map[int]int)

	for i := range score {
		mapRes[score[i]] = i
	}

	sort.Ints(score)
	fmt.Println(score)
	x := 4

	for i := len(score) - 1; i >= 0; i-- {
		switch i {
		case len(score) - 1:
			res[mapRes[score[i]]] = "Gold Medal"
			break
		case len(score) - 2:
			res[mapRes[score[i]]] = "Silver Medal"
			break
		case len(score) - 3:
			res[mapRes[score[i]]] = "Bronze Medal"
			break
		default:
			res[mapRes[score[i]]] = strconv.Itoa(x)
			x++
			break

		}

	}

	return res
}

func main() {
	a := []int{4, 3, 5, 6, 2, 1}
	fmt.Println(findRelativeRanks(a))
}
