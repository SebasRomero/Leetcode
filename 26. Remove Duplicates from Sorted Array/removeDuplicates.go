package main

func removeDuplicates(list []int) int {
	count := 1
	for i := 1; i < len(list); i++ {
		if list[i] != list[i-1] {
			list[count] = list[i]
			count++
		}
	}
	return count
}
