package main

import (
	 "fmt"
)

func main() {

	nums:=[]int{
		1,2,3,4,5,6,7,
	}

	result:=rotate(nums,2)

	fmt.Printf("%d",result)
}

func rotate(nums []int,k int )[]int {
	count :=len(nums)

	for i := 0; i < k; k-- {  //循环看k值
		   heigh_math :=nums[count-1]//获取最高位的数组
		   for j := count-2; j >=0; j-- { //从倒数第二个数开始遍历数组,把所有下标的数组抬升一位
		   	   nums[j+1]=nums[j]
		   }
		   //最后把弹出的最高位放到最初的位置
		   nums[0]=heigh_math
	}


	return nums
}