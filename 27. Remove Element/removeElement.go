package main

func removeElement(list []int, num int) int {
	for i, val := range list {
		if val == num {
			list = append(list[:i], list[i+1:]...)
			return removeElement(list, num)
		}
	}
	return len(list)
}
