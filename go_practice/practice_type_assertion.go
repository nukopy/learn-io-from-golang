package main

import "fmt"

/*
Type Assertion により、構造体が特定のインタフェースを実装しているかチェックする。
コンパイル時、または実行時にチェックする方法がある。
*/

/* ------------------------------ */
/* compile time check */
/* ------------------------------ */

type pusher interface {
	push()
	addListener()
}

type pusherImpl struct {}

func (p pusherImpl) push() {}
func (p pusherImpl) addListener() {}

type pusherNotImpl struct {}
func (p pusherNotImpl) push() {}
// func (p pusherNotImpl) addListener() {}

// コンパイル時にインタフェースを実装しているかチェックする
var _ pusher = pusherImpl{} // ok
var _ pusher = (*pusherImpl)(nil) // ok
// var _ pusher = pusherNotImpl{} // コンパイルエラーになる
/*
cannot use pusherNotImpl{} (value of type pusherNotImpl) as pusher value in variable declaration: pusherNotImpl does not implement pusher (missing method addListener)
*/

/* ------------------------------ */
/* runtime check */
/* ------------------------------ */

func practiceNil() {
	fmt.Println("----- practice type assertion -----")

	p := (*pusherImpl)(nil)
	fmt.Printf("%T\n", p)

	practiceTypeAssertion()
	practiceTypeSwitch()
}

func practiceTypeAssertion() {
	// empty
	var v interface{}
	v = 100
	n, ok := v.(int)
	fmt.Printf("%T %v %v\n", n, n, ok)

	n1, ok1 := v.(string)
	fmt.Printf("%T %v %v\n", n1, n1, ok1)
}

func practiceTypeSwitch() {
	fmt.Println("--- practice type switch ---")
	var i interface{}
	i = 100

	switch v := i.(type) {
	case int:
		fmt.Printf("%T %v\n", v, v)
	case string:
		fmt.Printf("%T %v\n", v, v)
	default:
		fmt.Printf("unknown type\n")
	}
}