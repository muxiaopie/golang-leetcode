package main

import (
	"sort"
)

/**
四重循环加上剪枝AC
三重循环加上左右指针居然MLE
晕。。。
*/
func fourSum(nums []int, target int) [][]int {
	var ret [][]int
	sort.Ints(nums)
	val := [...]int{0, 0, 0}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			val[0] = nums[i] + nums[j]
			if val[0] > target && nums[j] > 0 { //后面不可能有解
				break
			}
			if j > i+1 && nums[j-1] == nums[j] { //相同元素
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				val[1] = val[0] + nums[k]
				if val[1] > target && nums[k] > 0 {
					break
				}
				if k > j+1 && nums[k-1] == nums[k] {
					continue
				}
				for m := k + 1; m < len(nums); m++ {
					val[2] = val[1] + nums[m]
					if val[2] == target {
						tmp := []int{nums[i], nums[j], nums[k], nums[m]}
						ret = append(ret, tmp)
						break
					}
					if val[2] > target {
						break
					}
				}
			}
		}
	}
	return ret
}
