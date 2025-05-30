package removeduplicates

import "sort"

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
