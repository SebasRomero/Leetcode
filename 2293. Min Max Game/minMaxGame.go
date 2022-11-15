package main

func minMaxGame(nums []int) int {
	newSlice := []int{}
	sw := true
	if len(nums) == 1 {
		return nums[0]
	} else {
		for i := 1; i < len(nums); i++ {
			if nums[i-1] == 1 || nums[i] == 1 {
				return 1
			} else {
				if i%2 != 0 && sw == true { // Odd
					if nums[i] < nums[i-1] {
						newSlice = append(newSlice, nums[i])
					} else {
						newSlice = append(newSlice, nums[i-1])
					}
					sw = false
				} else {
					if i%2 != 0 && sw == false { // Even
						if nums[i] > nums[i-1] {
							newSlice = append(newSlice, nums[i])
						} else {
							newSlice = append(newSlice, nums[i-1])
						}
						sw = true
					}
				}
			}
		}
	}
	return minMaxGame(newSlice)

}
