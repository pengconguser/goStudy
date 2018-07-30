package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	goss := runtime.GOOS

	fmt.Printf("system:%v \n", goss)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
