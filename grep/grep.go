package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	if len(os.Args) == 2{
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan(){
			if strings.Contains(scanner.Text(), os.Args[1]){
				fmt.Println(scanner.Text())
			}
		}
	}else if len(os.Args) == 3{
		file, err := os.Open(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), os.Args[1]){
				fmt.Println(scanner.Text())
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}