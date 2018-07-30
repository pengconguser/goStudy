package main

import ("fmt")

func main() {
    nums:=[]int{
    	1,2,3,4,5,6,7,
    }

	result:=rotate(nums,4)

	fmt.Printf("%d",result);
}


func rotate(nums []int, k int) {
	n := len(nums)

	if k = k % n; k <= 0 {
		return
	}

	_rotate(nums, 0, n-k, n)
}

func _rotate(nums []int, l, r, n int) {
	for i := 0; i < r-l && i < n-r; i++ {
		nums[l+i], nums[r+i] = nums[r+i], nums[l+i]
	}

	if r-l < n-r {
		_rotate(nums, r, r+(r-l), n)
	} else if r-l > n-r {
		_rotate(nums, l+(n-r), r, n)
	}
}