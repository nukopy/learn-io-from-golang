package practice

import "fmt"

/* ------------------------------------ */
/* 構造体へのインタフェース埋め込みの練習 */
// ref: Goで構造体へのインタフェース埋め込み活用例（aws-sdk-goの事例など） https://horizoon.jp/post/2019/03/16/go_embedded_interface/
// 1. 一部メソッドの別処理に差し替える
// 2. 特定のメソッドのみを実装する
/* ------------------------------------ */

type MyInterface interface {
	MethodX()
	MethodY()
}

// define struct A
// MyInterface を実装する
type A struct {
	Str string
}

func (a *A) MethodX() {
	msg := fmt.Sprintf("%s: MethodX", a.Str)
	fmt.Println(msg)
}

func (a *A) MethodY() {
	msg := fmt.Sprintf("%s: MethodY", a.Str)
	fmt.Println(msg)
}

// define struct B
// MyInterface を実装する
type B struct {
	// 無名インタフェースの埋め込み
	MyInterface
}

// func (b *B) MethodX() {
//     fmt.Println("B: MethodX")
// }

func (b *B) MethodY() {
	fmt.Println("Replaced MethodY")
}

func ReplaceMethodY(org MyInterface) MyInterface {
	return &B{org}
}

func PracticeInterfaceEmbedding() {
	a := &A{Str: "I am A"}
	a.MethodX() // I am A: MethodX
	a.MethodY() // I am A: MethodY

	b := ReplaceMethodY(a)
	b.MethodX() // I am A: MethodX
	b.MethodY() // Replaced MethodY

	bb := ReplaceMethodY(&B{a})
	bb.MethodX() // I am A: MethodX
	bb.MethodY() // Replaced MethodY
}

// 2. 特定のメソッドのみを実装する
type MyIntf interface {
	GetStrVal() string
	GetIntVal() int
}

// Intfを実装するHoge1構造体
// 2メソッド漏れなく実装する
type SimpleStruct struct {
	StrVal string
	IntVal int
}

func (s *SimpleStruct) GetStrVal() string {
	return s.StrVal
}

func (s *SimpleStruct) GetIntVal() int {
	return s.IntVal
}

// 無名インタフェース Intf を埋め込んだHoge2構造体
// GetIntVal メソッドのみを実装する
type SimpleStruct2 struct {
	MyIntf
}

func (s *SimpleStruct2) GetIntVal() int {
	return 100
}

// interfaceの実装済チェック
// SimpleStruct2 は GetStrVal メソッドを実装していないがコンパイルエラーにならない
var _ MyIntf = (*SimpleStruct)(nil)
var _ MyIntf = (*SimpleStruct2)(nil)

func PracticeInterfaceEmbedding2() {
	ss := new(SimpleStruct)
	fmt.Println(ss.GetStrVal()) // ""
	fmt.Println(ss.GetIntVal()) // 0

	ss2 := new(SimpleStruct2)
	fmt.Println(ss2.GetIntVal()) // 100
	// fmt.Println(ss2.GetStrVal()) // panic: runtime error: invalid memory address or nil pointer dereference
	// GetStrVal メソッドは SimpleStruct2 では実装されていないので、呼び出そうとすると panic が発生する
}
