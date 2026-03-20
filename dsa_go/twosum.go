


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




func numberOfSubmatrices(grid [][]byte) int {
    m := len(grid)
    n := len(grid[0])

    balance := make([][]int, m)
    countx := make([][]int, m)

    for i := 0; i < m; i++ {
        balance[i] = make([]int, n)
        countx[i] = make([]int, n)
    }

    count := 0

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {

            val := 0
            xval := 0

            if grid[i][j] == 'X' {
                val = 1
                xval = 1
            } else if grid[i][j] == 'Y' {
                val = -1
            }

            balance[i][j] = val
            countx[i][j] = xval

            if i > 0 {
                balance[i][j] += balance[i-1][j]
                countx[i][j] += countx[i-1][j]
            }
            if j > 0 {
                balance[i][j] += balance[i][j-1]
                countx[i][j] += countx[i][j-1]
            }
            if i > 0 && j > 0 {
                balance[i][j] -= balance[i-1][j-1]
                countx[i][j] -= countx[i-1][j-1]
            }

            if balance[i][j] == 0 && countx[i][j] > 0 {
                count++
            }
        }
    }

    return count
}
