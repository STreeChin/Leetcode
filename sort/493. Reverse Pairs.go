/*
Given an array nums, we call (i, j) an important reverse pair if i < j and nums[i] > 2*nums[j].
You need to return the number of important reverse pairs in the given array.
Example1: Input: [1,3,2,3,1] Output: 2
Example2: Input: [2,4,3,5,1] Output: 3
*/
/*（1）solution
Like merge sort, after the left and right intervals are ordered, 
count the number of elements that meet the conditions, then merge
*/

func reversePairs(nums []int) int {
    if len(nums) == 0 { return 0}
    return InversePairsCore(nums, 0, len(nums)-1)
}
func InversePairsCore(data []int, start, end int) int {
	if start == end { return 0 }
	mid := (start + end) >> 1
	lCount := InversePairsCore(data, start, mid) 
	rCount := InversePairsCore(data, mid+1, end)
	curCount, i, j := 0, start, mid+1
	for ; i <= mid && j <= end; {
		if data[i] > data[j]*2 {
			j++
			curCount += mid - i + 1 //Elements from i to mid reverse with the j
		} else {
			i++
		}
	}
	merge(data, start, mid, end)
	return lCount + rCount + curCount
}
func merge(data []int, start, mid, end int) {
	L, R := make([]int, mid-start+1), make([]int, end-mid)
	for i, v := range data[start : mid+1] {
		L[i] = v
	}
	for i, v := range data[mid+1 : end+1] {
		R[i] = v
	}
	for i, j, k := 0, 0, start; k <= end; k++ {
		if j >= len(R) || (i < len(L) && L[i] <= R[j]) {
			data[k] = L[i]
			i++
		} else {
			data[k] = R[j]
			j++
		}
	}
}
