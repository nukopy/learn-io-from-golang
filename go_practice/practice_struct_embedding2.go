package main

import (
	"fmt"
)

type baseStruct struct {
	num int
}

func (b baseStruct) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	baseStruct // 構造体の（無名）埋め込み
	str        string
}

type describer interface {
	describe() string
}

func PracticeStructEmbedding3() {
	// 初期化
	co := container{
		baseStruct: baseStruct{
			num: 1,
		},
		str: "some name",
	}

	// 埋め込んだ baseStruct 構造体のフィールドには co から直接アクセスできる
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// co から baseStruct のフィールドを呼び出すこともできる
	fmt.Println("also num:", co.baseStruct.num)

	// 埋め込んだ baseStruct 構造体のメソッドは直接呼び出せる
	fmt.Println("describe:", co.describe())

	// co から baseStruct のメソッドを呼び出すこともできる
	// co.baseStruct.describe() は co.describe() のシンタックスシュガー（糖衣構文）である
	// 継承ではなく委譲：継承とは異なり、埋め込みはコンパイル時に展開される
	fmt.Println("also describe:", co.baseStruct.describe())

	// Embedding structs with methods may be used to bestow interface implementations onto other structs. Here we see that a container now implements the describer interface because it embeds base.
	// メソッドを持つ構造体を埋め込むことで、他の構造体にインタフェースの実装を与えることができる。
	// ここでは、container が、describe メソッドを実装した baseStruct を埋め込むことで、
	// container に describer インターフェースを実装することができることを確認している。
	var d describer = co
	fmt.Println("describer:", d.describe())
}
