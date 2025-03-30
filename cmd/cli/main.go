package main

import (
	"fmt"
	mlib "github.com/rf742/zapper/internal/lib"
	"os"
)

func main() {
	data := mlib.ReadCsvFile(os.Args[1])
	for i:=0; i<len(data); i++ {
		fmt.Println(data[i])
	}
}
