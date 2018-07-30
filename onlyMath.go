package main

import "fmt"

func main() {
    nums :=[]int{
       4,1,2,1,2,
    }

    result:=singleNumber(nums)

    fmt.Printf("%d",result)
}


func singleNumber(nums []int) int {
    count :=len(nums)
    result:=0
    //大神的代码 异或过滤
    for i := 0; i < count; i++ {
         result ^= nums[i]
    }
    
    return result
}
