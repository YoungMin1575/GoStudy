package main

import (
	"os"
	"io"
	"fmt"
)

func main() {
	if len(os.Args) > 1 {

		srcPath := os.Args[1]
		destPath := os.Args[2]

		srcFile, err := os.Open(srcPath)
		defer srcFile.Close()
		check(err)

		destFile, err := os.Create(destPath)
		defer destFile.Close()
		check(err)

		_, err = io.Copy(destFile, srcFile)
		check(err)

		err = destFile.Sync()
		check(err)
	}
}

func check(err error) {
	if err != nil {
		fmt.Println("Error : %s", err.Error())
		os.Exit(1)
	}
}
