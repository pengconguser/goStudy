package main

import (
	 "fmt"
)

func main() {

	nums:=[]int{
		3,2,4,
	}

	target:= 6
	result:=twoSum(nums,target)

	fmt.Printf("%v \n",result)
}

func twoSum(nums []int,target int)[]int {
	 result :=make([]int,0)
	 count  :=len(nums)

	 for i,value := range nums{
	   for j := i+1; j <=count-1 ; j++ {
			if(value+nums[j]==target && i!=j){                         
	 	 	    result=append(result,i)
	 	 	    result=append(result,j)
	 	 	    break;
	 		}
	   }
	 }
 
	 return result
}  