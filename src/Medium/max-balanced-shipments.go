// https://leetcode.com/contest/weekly-contest-461/problems/maximum-balanced-shipments/description/

package medium

import (
	stk "leetcode/collections/stack"
)

func maxBalancedShipments(weight []int) int {
	// for each parcel, it would be useful to track the index of the previous larger parcel
	//  0,  1, 2, 3, 4, 5
	// -1, -1, 1, 1, 3, 4
	// greedy approach after getting the previous larger parcel index
	// track the parcels that have so far been added to the shipment
	n := len(weight)
	prevStack := stk.NewStack()
	prevLarger := make([]int, n)
	for idx, val := range weight {
		// pop the stack till the element is smaller than it
		for {
			top := prevStack.Peek()
			if top == nil {
				prevLarger[idx] = -1
				break
			}
			topElem := top.(int)
			if weight[topElem] <= val {
				prevStack.Pop()
			} else {
				prevLarger[idx] = topElem
				break
			}
		}
		prevStack.Push(idx)
	}
	//fmt.Println(prevLarger)

	// track the parcels that have so far been added to the shipment
	lastAdded := -1
	countShipments := 0
	for currIdx, prevIdx := range prevLarger {
		if prevIdx == -1 {
			continue
		}
		// add from prevIdx to currIdx
		if lastAdded >= prevIdx {
			// can't add the current subarray to shipments as some packages from it already added
			continue
		}
		lastAdded = currIdx
		countShipments++
	}
	return countShipments
}
