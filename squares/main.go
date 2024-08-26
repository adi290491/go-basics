package main

import "fmt"

// func square() func() int {
// 	var x int
// 	return func() int {
// 		x++
// 		return x * x
// 	}
// }

func main() {

	// f := square()
	// g := cube()

	// var h func(int) func() float64

	// h = func(mul int) func() float64 {
	// 	var x int
	// 	return func() float64 {
	// 		x++
	// 		x *= mul
	// 		return math.Sqrt(float64(x))
	// 	}
	// }
	// fn := h(2)

	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("Square: %*s%d\n", 4, "*", f())
	// 	fmt.Printf("Cube: %*s%d\n", 8, "*", g())
	// 	fmt.Printf("Squareroot: %*s%.2f\n", 16, "*", fn())
	// }
	square([]int{1, 2, 3})
	for _, c := range cubes {
		c()
	}
	a := []int{}

	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 190))
	fmt.Println(sum(a...))

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s2 = append(s2, s1...)
	fmt.Println(sum(s2...))
	err := s1[100]
	panic(err)
}

func sum(vals ...int) (total int) {

	for _, val := range vals {
		total += val
	}
	return
}

var cubes []func()

func square(ints []int) {

	for _, i := range ints {
		cubes = append(cubes, func() {
			cube(i)
		})
	}
}

func cube(x int) {
	fmt.Println(x * x * x)
}
