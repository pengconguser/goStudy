package main

import (
	 "fmt"
)

func main() {
	 nums :=[]int{
	 	0,1,0,3,12,
	 }

	 result:=moveZero(nums)

	 fmt.Printf("%v",result)
}

func moveZero(nums []int )[]int {
	 count :=len(nums)

	 for i := count-1; i >= 0; i-- {
	 	  if(nums[i]==0){
	 	  	 nums=append(nums[:i],nums[i+1:]...)
	 	  	 nums=append(nums,0)
	 	  }
	 }

	 return nums
}