


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


func incdec(arr int[]) int[]{
	n:=len(arr)
	sort.Ints(arr)
	mid:=n/2
	left:=mid
	right:=n-1

	for left<right{
		arr[left],arr[right]= arr[right],arr[left]
		left++
		right--
	}
	return arr
}
func bitwiseComplement(n int) int{
	if n==0{
		return 1}
	mask:=1
	for mask <=n{
		mask<<=1
		}
	return (mask-1)^n}

