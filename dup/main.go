package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	for {
		line := in.Text()

		// if err == io.EOF {
		// 	break
		// }
		counts[string(line)]++
	}

	fmt.Println(counts)

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
