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


func duplicate3func (nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}

	buckets := make(map[int64]int64)
	width := int64(t) + 1

	for i, num := range nums {
		n := int64(num)
		bucketID := getBucketID(n, width)

		// 1️⃣ Same bucket
		if _, exists := buckets[bucketID]; exists {
			return true
		}

		// 2️⃣ Previous bucket
		if val, exists := buckets[bucketID-1]; exists && abs(n-val) <= int64(t) {
			return true
		}

		// 3️⃣ Next bucket
		if val, exists := buckets[bucketID+1]; exists && abs(n-val) <= int64(t) {
			return true
		}

		// Insert into bucket
		buckets[bucketID] = n

		// Maintain sliding window
		if i >= k {
			old := int64(nums[i-k])
			delete(buckets, getBucketID(old, width))
		}
	}

	return false
}

func getBucketID(num, width int64) int64 {
	if num >= 0 {
		return num / width
	}
	return (num+1)/width - 1
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}