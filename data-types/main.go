package main

import (
	"fmt"
	"slices"
	"unicode"
)

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Printf("%T %T", s[0], s[7])
	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])
	fmt.Println(s[:5] + " goodbye")
	hello := s[:5]
	world := s[7:]
	fmt.Println(hello, &hello)
	fmt.Println(world, &world)
	fmt.Println(s, &s)

	for _, c := range s {
		fmt.Printf("%T\n", c)
	}
	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Printf("%T\n", s[i])
	}
	n := 0
	for range s {
		n++
	}

	fmt.Println(n)
	r := []rune(s)

	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-i-1] = r[len(r)-i-1], r[i]
	}

	fmt.Println(string(r))

	r = []rune("HeLlO, WorLD")
	for i := 0; i < len(r); i++ {
		if unicode.IsUpper(r[i]) {
			r[i] = unicode.ToLower(r[i])
		} else {
			r[i] = unicode.ToUpper(r[i])
		}
	}

	fmt.Println(string(r))

	symbol := [...]int{0: 1, 2: 3, 1: 3}
	fmt.Println(symbol)

	for _, v := range symbol {
		fmt.Println(v)
	}

	fmt.Println(comma("12345678901"))

	a := []int{1, 2, 3, 4, 5} // [4, 5, 1, 2, 3]
	//slices.Reverse(a[:])
	rotate(a[:], 4)
	fmt.Println(a[1:3:4])

	b := []int{1, 7, 3, 6, 5, 4, 3}

	fmt.Println(b)
	fmt.Println(b[1:3:4])
	maxV := b[0]
	minV := b[0]
	for _, v := range b {
		maxV = max(maxV, v)
		minV = min(minV, v)
	}

	fmt.Println(maxV, minV)
	slices.Sort(b)
	fmt.Println(b)

	c := []string{"apple", "banana", "orange", "beans", "pineapple"}
	slices.Sort(c)
	fmt.Println(c)
	slices.SortFunc(c, func(c1, c2 string) int {
		return len(c2) - len(c1)
	})

	fmt.Println(c)

	ages := make(map[string]int)
	ages["john"] = 32
	ages["jimmy"] = 34
	fmt.Println(ages)
	delete(ages, "jimmy")
	fmt.Println(ages)
	fmt.Println(ages["carl"])

	Add(c)
	fmt.Printf("%q\n", m)
	fmt.Println(Count(c))

}

var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) { m[k(list)]++ }

func Count(list []string) int { return m[k(list)] }

func rotate(arr []int, k int) {
	k = len(arr) % k
	slices.Reverse(arr[len(arr)-k:])
	slices.Reverse(arr[:len(arr)-k])
	slices.Reverse(arr)

}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + ", " + s[n-3:]
}
