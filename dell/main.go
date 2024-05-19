package dell

// Удалить повторяющиеся элементы из массива. Сдвигать все элементы массива на место удаленных. В таком случае в конце массива должны получиться нули
func RemoveDuplicates(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	for k := i + 1; k < len(nums); k++ {
		nums[k] = 0
	}

	return nums
}
