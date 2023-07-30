package practice

import "fmt"

func PracticeSlicesLenCap() {
	sliceBasic()
	sliceWithMake()
	sliceAppend()
}

func sliceBasic() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	// Access idx = 0, 1
	s = s[:2]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

/* 出力
len=6 cap=6 [2 3 5 7 11 13]
len=0 cap=6 []
len=4 cap=6 [2 3 5 7]
len=2 cap=4 [5 7]
len=2 cap=4 [5 7]
*/

func sliceWithMake() {
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5) // make([]T, length, capacity)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

/* 出力
a len=5 cap=5 [0 0 0 0 0]
b len=0 cap=5 []
c len=2 cap=5 [0 0]
d len=3 cap=3 [0 0 0]
*/

func sliceAppend() {
	var s []int // ゼロ値は nil。nil スライスの len と cap は 0 である。
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4) // cap の割り当ては多めに行われることがある。詳しくは https://blog.golang.org/go-slices-usage-and-internals を参照。
	printSlice(s)
}

/*
len=0 cap=0 []
len=1 cap=1 [0]
len=2 cap=2 [0 1]
len=5 cap=6 [0 1 2 3 4]
*/