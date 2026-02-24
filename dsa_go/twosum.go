


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