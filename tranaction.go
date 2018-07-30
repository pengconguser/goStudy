package main

import "fmt"


func main() {

	prices:=[]int{
		7,1,5,3,6,4,
	}

	result:=maxProfit(prices)

	fmt.Printf("%d",result)
}

func maxProfit(prices [] int) int {
	 count :=len(prices)

	 total :=0;

	 for i := 0; i < count-1; i++ {
	 	  if prices[i+1]>prices[i] {
	 	     total+=prices[i+1]-prices[i]	
	 	  } 	
	 }

	 return total
}