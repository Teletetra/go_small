package main

func duplicate(nums []int)bool{
	seen:=make(map[int]struct{})
	for _,num:=range nums{
		if _,exists:=seen[num];exists{
			return true
		}
		seen[num]=struct{}{}
	}
	return false
}

func duplicate2(nums []int,k int)bool{
	seen:=make(map[int]int)

	for i,num:=range nums{
		if _,exists:=seen[num];exists{
				if (i-seen[num])<=k{
					return true
				}
		}	
	}
	return false
}