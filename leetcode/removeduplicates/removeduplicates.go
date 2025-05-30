package removeduplicates

import "sort"

// Algo complexity O(n), O(1) by mem.
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	a := 0
	b := 1
	for b < len(nums) {
		if nums[a] == nums[b] {
			b++
			continue
		}

		if (b - a) == 1 {
			a++
			b++
			continue
		}

		a++
		nums[a] = nums[b]
		b++
	}

	return a + 1
}

func trivialRemoveDuplicates(nums []int) int {
	unique := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := unique[v]; !ok {
			unique[v] = struct{}{}
			continue
		}
	}

	var res []int
	for k := range unique {
		res = append(res, k)
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })

	for n, v := range res {
		nums[n] = v
	}

	return len(res)
}
