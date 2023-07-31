package main

import (
	"fmt"
	"os"
)

func testFileIo() {
	testFileRead()
	testFileWrite()
}

func testFileRead() {
	// open file
	f, err := os.Open("./file_io_read.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// read file
	buf := make([]byte, 1024)
	count, err := f.Read(buf) // reader to buf
	if err != nil {
		fmt.Println(err)
		return
	}

	// output
	str := string(buf[:count])
	fmt.Printf("read %d bytes: %q\n", count, buf[:count])
	fmt.Println(str)

	/* 出力結果
	read 28 bytes: "Hello, world!\nHello, Golang!"
	Hello, world!
	Hello, Golang!
	*/
}

func testFileWrite() {
	// create file
	f, err := os.Create("./file_io_write.txt")
	// f, err := os.Open("./file_io_write.txt")
	/* os.Open だと書き込み権限がないため、以下のようなランタイムエラーが発生する
	error: write ./file_io_write.txt: bad file descriptor
	*/
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer f.Close()
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// write file
	str := "Write this file by Golang!\nHello, world!\nHello, Golang!"
	data := []byte(str)
	count, err := f.Write(data) // buf to writer
	if err != nil {
		fmt.Println(err)
		return
	}

	// output
	fmt.Printf("write %d bytes\n", count)

	/* 出力結果
	write 55 bytes
	*/
}