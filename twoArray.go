package main

import (
	"fmt"
)

func main() {

	num1:=[]int{
		1,2,2,1,
	}

	num2:=[]int{
		2,2,
	}


	result :=intersect(num1,num2)

	fmt.Printf("%d",result)
}

func intersect(num1[]int,num2[]int)[]int {
	count_num1:=len(num1)
	count_num2:=len(num2)

    //创建切片
	map_hash :=make([]int,0)
    result_slice :=make([]int,0)

    for i := 0; i < count_num1; i++ {
    	map_hash =append(num1[i])
    }

    for i := 0; i < count_num2; i++ {
    	if count_num2>0 {
    	 	
    	}
    }

}
