package main

import (
	"fmt"
)

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Initialize
	maxEndingHere := nums[0]
	minEndingHere := nums[0]
	globalMax := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		// If num is negative, swap max and min
		if num < 0 {
			maxEndingHere, minEndingHere = minEndingHere, maxEndingHere
		}

		// Update max/min ending here
		maxEndingHere = max(num, maxEndingHere*num)
		minEndingHere = min(num, minEndingHere*num)

		// Update global maximum
		if maxEndingHere > globalMax {
			globalMax = maxEndingHere
		}
	}

	return globalMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	arr := []int{2, 3, -2, 4}
	fmt.Println("Maximum Product Subarray:", maxProduct(arr)) // Output: 6
}

func lengthOfLongestSubstring(s string) int {
    maxLen:=0
    left:=0
    char_set:=make(map[rune]int)
    for right,ch:=range s{
        if prevIndex,exists:=char_set[ch];exists {
            if prevIndex>=left{
                left=prevIndex+1
            }
        }
        char_set[ch]=right
        if right-left+1>maxLen{
            maxLen=right-left+1
        }
    }
    return maxLen
}