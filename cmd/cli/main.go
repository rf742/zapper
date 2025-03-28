package main

import(
	"fmt"
	mlib "github.com/rf742/zapper/internal/lib"
	"os"
)

func main() {
	fmt.Printf("The charge on an electron is %e\n", mlib.E)
	data := mlib.ReadCsvFile(os.Args[1])
	for i:=0; i<len(data); i++ {
		fmt.Println(data[i])
	}
	p := mlib.Point{ 1.0, 2.0, 4.0 }
	fmt.Println(p.Charge())
	fmt.Println(p.Position())
}
