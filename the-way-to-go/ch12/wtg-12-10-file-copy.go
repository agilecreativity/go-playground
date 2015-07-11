package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)

	if err != nil {
		fmt.Printf("Error opening source file %s\n", srcName)
		return
	}

	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return
	}

	defer dst.Close()

	return io.Copy(dst, src)
}

func main() {
	_, err := CopyFile("target.txt", "source.txt")
	if err != nil {
		return
	}
	fmt.Println("Done copying file")
}
