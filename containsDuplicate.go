package main

import (
	 "fmt"
)

func main() {
	
	 nums:=[]int{
	 	1,2,3,4,4,5,6,7,8,
	 }

	result:=containsDumplicate(nums)

	fmt.Printf("%b",result)
}

func containsDumplicate(nums []int) bool {
	count :=len(nums)
	
	for i := 0; i < count; i++ {
		  for j := 0; j < count; j++ {
		  	  if (nums[i]==nums[j] && i!=j) {
		  	  	  return true
		  	  }
		  }
	}

	return false
}