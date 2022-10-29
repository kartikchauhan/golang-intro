package main

import (
	"fmt"
	"io"
	"os"
)

// srcFile & destFile will be closed even if os.OpenFile() or io.Copy() gives any error
func copyFile(dest string, src string) (written int, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()

	destFile, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return
		// destFile, _ = os.Create(dest)
	}
	defer destFile.Close()

	bytesCopied, err := io.Copy(destFile, srcFile)
	if err != nil {
		return
	}

	return int(bytesCopied), err
}

func main() {
	dest := os.Args[1]
	bytes, err := copyFile(dest+".txt", "srcFile.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		panic("Gojira is here!! Run!!")
	}

	fmt.Printf("[%v bytes]", bytes)

}
