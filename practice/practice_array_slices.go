package practice

import (
	"fmt"
)

func practiceArraySlices() {
	fmt.Println("----- Practice Array Slices -----")

	// array
	const len int = 10 // len は const でないといけない
	var arr [len]int // 配列はゼロ値で初期化される
	fmt.Println(arr)

	arr2 := [len]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	// var lenVar int = 10
	// var arrVar [lenVar]int --> コンパイルエラーになる: invalid array length lenVar
	// fmt.Println(arrVar)
}