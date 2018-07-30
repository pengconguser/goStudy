package main

import (
	"fmt"
)

func main() {
	digits:=[]int{
		1,2,3,
	}

	result:=removeDigits(digits)

	fmt.Printf("%d \n",result)
}


func removeDigits(digits []int)[]int {
	 count:=len(digits)

	 for i := count-1; i >=0 ;i-- {
	 	 if(digits[i]<9){
	 	 	 digits[i]++
	 	 	 return digits
	 	 }

	 	 digits[i] = 0
	 }

	 res:=make([]int,count+1)

	 res[0]=1

	 return res
}  