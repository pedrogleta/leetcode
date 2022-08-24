package main

func canCompleteCircuit(gas []int, cost []int) int {
	totalSurplus := 0
	surplus := 0
	start := 0

	for i, thisGas := range gas {
		thisCost := cost[i]
		totalSurplus += thisGas - thisCost
		surplus += thisGas - thisCost
		if surplus < 0 {
			surplus = 0
			start = i + 1
		}
	}

	if totalSurplus < 0 {
		return -1
	}

	return start
}
