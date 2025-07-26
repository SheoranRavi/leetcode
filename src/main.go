// this file is for testing the solutions locally
package main

import (
	"fmt"
	"leetcode/medium"
)

func main() {
	edges := [][]int{{6, 2, 35129}, {3, 4, 99499}, {2, 7, 43547}, {8, 1, 78671}, {2, 1, 66308}, {9, 6, 33462}, {5, 1, 48249}, {2, 3, 44414}, {6, 7, 44602}, {1, 7, 14931}, {8, 9, 38171}, {4, 5, 30827}, {3, 9, 79166}, {4, 8, 93731}, {5, 9, 64068}, {7, 5, 17741}, {6, 3, 76017}, {9, 4, 72244}}
	numPaths := medium.CountRestrictedPaths(9, edges)
	fmt.Println("Number of restricted paths:", numPaths)
}
