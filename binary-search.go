package main

func binary(nums []int, target int) int {
	var down = 0
	var up = len(nums) - 1

	for down <= up {
		var middle = (up + down) / 2
		var guess = nums[middle]
		if guess == target {
			return middle
		}
		if guess > target {
			up = middle - 1
		} else {
			down = middle + 1
		}
	}
	return -1
}
