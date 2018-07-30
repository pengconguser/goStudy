package main

import "fmt"

func main() {
	
	nums :=[]int{1,1,2}

	result:=removeDuplicates(nums)

	fmt.Printf("%d",result)
}

func removeDuplicates(nums []int)int {
    if len(nums)==0 {
    	return 0
    }

    number :=0 //标记数组计数
    count:=len(nums)

    for i := 0; i < count; i++ {
     	 if(nums[i]!=nums[number]){
     	 	 number++
     	 	 nums[number]=nums[i]
     	 }
    }

    number+=1

    return number
}