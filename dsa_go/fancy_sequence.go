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
