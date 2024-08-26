package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}

	buf.WriteByte(']')
	return buf.String()
}

func main() {
	// fmt.Println(intsToString([]int{1, 2, 3}))

	// fmt.Println(comma(os.Args[1]))
	// fmt.Println(comma2(os.Args[1]))

	fmt.Println(anagram(os.Args[1], os.Args[2]))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {

	parts := strings.Split(s, ".")

	integerPart := parts[0]
	fractionPart := ""
	if len(parts) > 1 {
		fractionPart = "." + parts[1]
	}

	var buf bytes.Buffer
	for i, v := range integerPart {
		if i > 0 && i%3 == 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%s", string(v))
	}
	return buf.String() + fractionPart
}

func anagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	chArr := [26]int{}

	for _, ch := range s {
		chArr[ch-'a']++
	}

	for _, ch := range t {
		chArr[ch-'a']--
		if chArr[ch-'a'] < 0 {
			return false
		}
	}

	return true
}
