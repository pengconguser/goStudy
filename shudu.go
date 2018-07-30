package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := [][]string{
		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}

	result := check_shudu(nums)

	fmt.Printf("%v", result)
}

func check_shudu(nums [][]string) bool {

	for i := 0; i < 9; i++ {
		bit_map_row := make([]int, 0)
		bit_map_col := make([]int, 0)
		bit_map_cube := make([]int, 0)
		//下标从1开始d
		for j := 0; j < 9; j++ {
			if nums[i][j] != "." {
				bit, _ := strconv.Atoi(nums[i][j])
				if bit_map_row[bit-1] == 1 {
					return false
				} else {
					bit_map_row[bit-1] = 1
				}
			}

			if nums[j][i] != "." {
				bit, _ := strconv.Atoi(nums[j][i])
				if bit_map_col[bit-1] == 1 {
					return false
				} else {
					bit_map_col[bit-1] = 1
				}
			}

			RowIndex := 3*(i/3) + j/3
			ColIndex := 3*(i%3) + j%3
			val := nums[RowIndex][ColIndex]

			if val != "." {
				val, _ := strconv.Atoi(val)
				if bit_map_cube[val-1] == 1 {
					return false
				} else {
					bit_map_cube[val-1] = 1
				}
			}
		}
	}

	return true
}
