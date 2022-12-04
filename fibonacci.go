package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var sequence []int
	sequence = append(sequence, 0)
	return func() int {
		last := sequence[len(sequence)-1]
		var secondLast int
		fmt.Println(sequence)

		if len(sequence) <= 1 {
			secondLast = 1
		} else {
			secondLast = sequence[len(sequence)-2]
		}
		sequence = append(sequence, last+secondLast)
		return sequence[len(sequence)-2]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
