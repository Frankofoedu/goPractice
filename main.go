package main

import (
	"fmt"
	"math"
)


func main() {
	c := []int32{8, 2, 4, 6}
	segment(2, c)
}

func segment(x int, space []int32) {
	var segment [][]int32

	for i := 0; i < len(space); i++ {
		if i == len(space)-1 {
			break
		}

		segment[i] = []int32{space[i], space[i+1]}
	
		//segment = append(segment,...)
	}
	fmt.Println(segment, "\n")
}

func findMinSubArray(S int32, arr []int32 ) int {
	// TODO: Write your code here

	var arraySum int32
	var windowStart = 0
	var minLength = math.MaxInt32
	for i := int32(0); i < int32(len(arr)); i++ {
		arraySum += arr[i]

		for i >= S {
			arraySum -= arr[windowStart]
			minLength = int(math.Min(float64(minLength), float64(i) - float64(windowStart + 1)))

			set
		}

	}
    return -1;
  }