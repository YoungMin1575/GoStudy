package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"time"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	file.Seek(0,os.SEEK_END)

	for  {
		reader := bufio.NewReader(file)
		strAppended, _ := Readln(reader)
		if strAppended != ""{
			fmt.Println(strAppended)
		}
		strAppended,_ = Readln(reader)
		time.Sleep(time.Microsecond)
	}
}

func Readln(reader *bufio.Reader) (string, error) {
	var (isPrefix bool = true
		err error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = reader.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln),err
}
