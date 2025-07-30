package removeelements

func removeElement(nums []int, val int) int {
	var k int
	for i := 0; i < len(nums); i++ {
		nums[k] = nums[i]

		if nums[i] != val {
			k++
		}
	}
	return k
}

func removeElementTrivial(nums []int, val int) int {
	i := 0
	for i < len(nums) {
		if nums[i] != val {
			i++
			continue
		}

		nums = append(nums[:i], nums[i+1:]...)
	}

	return len(nums)
}
