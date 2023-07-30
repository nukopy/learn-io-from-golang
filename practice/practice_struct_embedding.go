// ref: https://qiita.com/momotaro98/items/4f6e2facc40a3f37c3c3
package practice

import (
	"fmt"
)

type Flyer interface {
	Fly() string
}

type Runner interface {
	Run() string
}

// ①  インターフェースにインターフェースを埋め込む
type FlyingRunner interface {
	Flyer
	Runner
}

// ② 構造体にインターフェースを埋め込む
type ToriJin struct { // 鳥人
	// 無名インターフェースの埋め込み
	FlyingRunner
}

// ③ 構造体に構造体を埋め込む
type ShinJinrui struct { // 新人類
	// ToriJin
	*ToriJin
}

// FlyingRunnerインターフェースを実装する型
type RealToriJin struct{}

func (r RealToriJin) Fly() string { return "Fly!" }

func (r RealToriJin) Run() string { return "Run!" }

func PracticeStructEmbedding() {
	aRealToriJin := &RealToriJin{}
	// ② 構造体ToriJinにFlyingRunnerインターフェースを
	// 実装しているRealToriJinの変数を埋め込む。
	aToriJin := &ToriJin{
		FlyingRunner: aRealToriJin,
	}
	// ③ 構造体ShinJinruiに構造体Torijinの変数を埋め込む。
	aShinJinrui := &ShinJinrui{
		// ToriJin: aToriJin,
		aToriJin,
	}
	fmt.Println(aShinJinrui.Fly()) // Fly!
	fmt.Println(aShinJinrui.Run()) // Run!
}

/* ------------------------------------ */
/* 構造体への構造体埋め込みの練習 */
/* ------------------------------------ */

type InnerStruct struct {
	Value int
}

type OuterStructByValue struct {
	InnerStruct
}

type OuterStructByPointer struct {
	*InnerStruct
}

func PracticeStructEmbedding2() {
	// inner 構造体の値を 42 に設定して初期化
	inner := InnerStruct{Value: 42}

	// inner 構造体、inner 構造体のポインタを埋め込んだ outer 構造体を初期化
	outerByValue := OuterStructByValue{InnerStruct: inner}
	outerByPointer := OuterStructByPointer{InnerStruct: &inner}

	// inner 構造体の値を変更
	inner.Value = 100

	// inner 構造体の値を埋め込んだ outer.InnerStruct の値は変更されず、inner 構造体のポインタ（参照）を埋め込んだ outer.InnerStruct の値は変更される
	fmt.Println(outerByValue.InnerStruct.Value)   // Output: 42
	fmt.Println(outerByPointer.InnerStruct.Value) // Output: 100
}
