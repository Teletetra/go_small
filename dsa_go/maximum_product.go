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


func rotateK(s string, k int) string{
	runes:=[]rune(s)
	n:=len(runes)
	if n==0{
		return s
	}

	k=k%n
	return string(runes[k:])+string(runes[:k])
}


func rotateSlice[T any](s []T,k int)[] T{
	n:=len(s)
	if n==0{
		return s
	}
	k=k%n
	return append(s[k:],s[:k]...)

}


func removeComments(source []string) []string {
    var result []string
    inBlock := false
    newline := ""

    for _, line := range source {
        i := 0

        if !inBlock {
            newline = ""
        }

        for i < len(line) {
            if !inBlock && i+1 < len(line) && line[i] == '/' && line[i+1] == '*' {
                inBlock = true
                i += 2
            } else if inBlock && i+1 < len(line) && line[i] == '*' && line[i+1] == '/' {
                inBlock = false
                i += 2
            } else if !inBlock && i+1 < len(line) && line[i] == '/' && line[i+1] == '/' {
                break
            } else if !inBlock {
                newline += string(line[i])
                i++
            } else {
                i++
            }
        }

        if !inBlock && len(newline) > 0 {
            result = append(result, newline)
        }
    }

    return result
}


type Fuck struct {
	name   string
	status string
}

func fuckme() Fuck {
	f := Fuck{
		name:   "kanishk",
		status: "doomed",
	}

	return f
}