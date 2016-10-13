package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) > 1 {

		srcPath := os.Args[1]
		destPath := os.Args[2]

		err := os.Rename(srcPath, destPath)

		if err != nil {
			fmt.Println(err)
		}
	}
}
