package intset

// precomputed table: pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	// fmt.Println(pc)
}

// Popcounts returns the population count (number of set bits) of x
func PopCount(x uint64) (ret int) {
	for i := 0; i < 8; i++ {
		ret += int(pc[byte(x>>(i*8))])
	}
	return
}

func PopCountShift(x uint64) (ret int) {
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			ret++
		}
		x >>= 1
	}
	return ret
}

func PopCountK(x uint64) (ret int) {
	for x != 0 {
		x &= (x - 1)
		ret++
	}
	return ret
}
