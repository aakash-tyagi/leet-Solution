package main

func findDuplicate(num []int) int {

	mapArr := make(map[int]int)

	for _, val := range num {
		mapArr[val]++
	}

	for k, val := range mapArr {
		if val > 1 {
			return k
		}
	}
	return 0
}

func main() {
	a := []int{1, 2, 3, 3, 4, 5, 6, 5, 9}
	findDuplicate(a)

}
