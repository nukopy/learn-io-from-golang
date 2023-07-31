package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

// ファイルの中身を読み込んで別のファイルへの書き込みを行う関数
func FileToFileTranslateIntoGerman(src *os.File, dst *os.File) {
	// reader -> buf
	buf := make([]byte, 300)
	n, err := src.Read(buf)
	if err != nil {
		panic(err)
	}
	log.Printf("%d bytes read", n)

	str := string(buf[:n])
	str = strings.ReplaceAll(str, "Hello", "Guten Tag")
	log.Println(str)

	// buf -> writer
	n, err = dst.Write(buf[:n])
	if err != nil {
		panic(err)
	}
	log.Printf("%d bytes written", n)
}

// ファイルの中身を読み込んでネットワークコネクションへ書き込む関数
func FileToNetworkConnTranslateIntoGerman(src *os.File, dst net.Conn) {
	// reader -> buf
	buf := make([]byte, 300)
	n, err := src.Read(buf)
	if err != nil {
		panic(err)
	}
	log.Printf("%d bytes read", n)

	str := string(buf[:n])
	str = strings.ReplaceAll(str, "Hello", "Guten Tag")
	log.Println(str)

	// buf -> writer
	n, err = dst.Write(buf[:n])
	if err != nil {
		panic(err)
	}
	log.Printf("%d bytes written", n)
}

// ネットワークコネクションから読み込んでファイルへ書き込む関数
func NetworkConnToFileTranslateIntoGerman(src net.Conn, dst *os.File) {
	// reader -> buf
	buf := make([]byte, 300)
	n, err := src.Read(buf)
	if err != nil {
		panic(err)
	}
	log.Printf("%d bytes read", n)

	str := string(buf[:n])
	str = strings.ReplaceAll(str, "Hello", "Guten Tag")
	log.Println(str)

	// buf -> writer
	n, err = dst.Write(buf[:n])
	if err != nil {
		panic(err)
	}
	log.Printf("%d bytes written", n)
}

// ネットワークコネクションから読み込んで別のネットワークコネクションへ書き込む関数
func NetworkConnToNetworkConnTranslateIntoGerman(src net.Conn, dst net.Conn) {
	// reader -> buf
	buf := make([]byte, 300)
	n, err := src.Read(buf)
	if err != nil {
		panic(err)
	}
	log.Printf("%d bytes read", n)

	str := string(buf[:n])
	str = strings.ReplaceAll(str, "Hello", "Guten Tag")
	log.Println(str)

	// buf -> writer
	n, err = dst.Write(buf[:n])
	if err != nil {
		panic(err)
	}
	log.Printf("%d bytes written", n)
}

// ↑の関数を抽象化した関数
// io.Reader、io.Writer を使えばこの関数だけでさまざまな入出力を扱えるようになる
func TranslateIntoGerman(src io.Reader, dst io.Writer) {
	// reader -> buf
	buf := make([]byte, 300)
	n, err := src.Read(buf)
	if err != nil {
		panic(err)
	}
	log.Printf("%d bytes read", n)

	str := string(buf[:n])
	str = strings.ReplaceAll(str, "Hello", "Guten Tag")
	log.Println(str)

	// buf -> writer
	n, err = dst.Write(buf[:n])
	if err != nil {
		panic(err)
	}
	log.Printf("%d bytes written", n)
}

func main() {
	fmt.Println("Hello, World!")
}