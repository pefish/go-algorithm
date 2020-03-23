package go_algorithm

import "fmt"

func ExampleShuffle() {
	a := []int{3,5,6,5,9,2}
	b := Shuffle(a)
	fmt.Println(a)
	fmt.Println(len(b))
	// Output:
	// [3 5 6 5 9 2]
	// 6
}