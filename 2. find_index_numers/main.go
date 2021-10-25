package main
import "fmt"

func twoSum(nums []int, target int) []int {
	 res := []int{}
	for i, v1 := range nums {
		news := nums[i+1:]
		for j, v2 := range news{
			if v1+v2 == target {
				res = append(res, i, j +(i+1) )
				return res
			}
		}
	}
	return res
}

func main() {
	nums := []int{3, 5, 4, 9}
	target := 12
	result := twoSum(nums, target)
	fmt.Println(result)
}
