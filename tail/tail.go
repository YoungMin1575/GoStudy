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

	/*scanner := bufio.NewScanner(file)
	lineCount := 1;

	for scanner.Scan() {
		lineCount++;
	}

	beginLineNum := lineCount - 5
	i := 1

	file.Seek(0,os.SEEK_SET)

	scanner2 := bufio.NewScanner(file)

	for scanner2.Scan() {
		if i >= beginLineNum {
			fmt.Println(scanner2.Text())
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}*/

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
