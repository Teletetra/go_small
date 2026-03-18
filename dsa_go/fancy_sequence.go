package main

const MOD int64 = 1000000007

type Fancy struct {
	seq []int64
	a   int64
	b   int64
}

func Constructor() Fancy {
	return Fancy{
		seq: []int64{},
		a:   1,
		b:   0,
	}
}

func modPow(x, n int64) int64 {
	res := int64(1)
	x %= MOD
	for n > 0 {
		if n%2 == 1 {
			res = (res * x) % MOD
		}
		x = (x * x) % MOD
		n /= 2
	}
	return res
}

func modInv(x int64) int64 {
	return modPow(x, MOD-2)
}

func (this *Fancy) Append(val int) {
	v := (int64(val) - this.b + MOD) % MOD
	v = (v * modInv(this.a)) % MOD
	this.seq = append(this.seq, v)
}

func (this *Fancy) AddAll(inc int) {
	this.b = (this.b + int64(inc)) % MOD
}

func (this *Fancy) MultAll(m int) {
	this.a = (this.a * int64(m)) % MOD
	this.b = (this.b * int64(m)) % MOD
}

func (this *Fancy) GetIndex(idx int) int {
	if idx >= len(this.seq) {
		return -1
	}
	return int((this.seq[idx]*this.a + this.b) % MOD)
}


func CountSubmetrices(grid [][]int,k int) int{
	m:=len(grid)
	n:=len(grid[0])
	count:=0
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if i>0{
				grid[i][j]+=grid[i-1][j]
				}
			if j>0{
				grid[i][j]+=grid[i][j-1]
				}
			if i>0 && j>0{
				grid[i][j]-=grid[i-1][j-1]
				}
			if grid[i][j]<=k{
				count++
				}
			}
		}
	return count
}



