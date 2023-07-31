package practice

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/* ref: https://go-tour-jp.appspot.com/methods/23
Exercise: rot13Reader
よくあるパターンは、別の io.Reader をラップし、ストリームの内容を何らかの方法で変換するio.Readerです。

例えば、 gzip.NewReader は、 io.Reader (gzipされたデータストリーム)を引数で受け取り、 *gzip.Reader を返します。 その *gzip.Reader は、 io.Reader (展開したデータストリーム)を実装しています。

io.Reader を実装し、 io.Reader でROT13 換字式暗号( substitution cipher )をすべてのアルファベットの文字に適用して読み出すように rot13Reader を実装してみてください。

rot13Reader 型は提供済みです。 この Read メソッドを実装することで io.Reader インタフェースを満たしてください。
*/

// io.Reader をラップして、ストリームの内容を変換する io.Reader
// rot13Reader は io.Reader をフィールドに持ち、自身も io.Reader を実装している。
// そのため、他の io.Reader をラップして、ストリームの内容を変換する io.Reader として利用できる。
type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(p []byte) (int, error) {
	n, err := rot.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}

	return n, err
}

func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'Z':
		return 'A' + (b - 'A' + 13) % 26
	case 'a' <= b && b <= 'z':
		return 'a' + (b - 'a' + 13) % 26
	default:
		return b
	}
}

func practiceIoReader() {
	fmt.Println("----- practice io.Reader -----")
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	// io.Copy: src から dst へコピーする
	dst := os.Stdout
	src := r
	// src := &r
	io.Copy(dst, src)
}
