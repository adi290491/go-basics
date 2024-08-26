package popcount

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var count byte = 0
	fmt.Println("Array: ", pc)
	for i := 0; i <= 7; i++ {
		fmt.Println("Count: ", count, x>>(i*8))
		count += pc[byte(x>>(i*8))]
	}

	return int(count)
}
