


func twoSum(nums []int, target int) []int {
    seen:=make(map[int]int)
    for i,num:=range nums{
        complement:=target-num
        if val,exists:=seen[complement];exists{
            return []int{val,i}
        }
        seen[num]=i
    }
    return nil
}


func freq(arr int[]) map[int]int{
	freq:= make(map[int]int)

	for _,num:=range arr{
		freq[num]++
	}
	return freq
}