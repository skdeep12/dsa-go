package array

func getConcatenation(nums []int) []int {
	nums = nums[:]
	nums = append(nums, nums...)
	return nums
}
