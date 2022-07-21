package scanner

import "fmt"

type MyScanner struct {
	foo int
}

func (mys *MyScanner) Scan(args ...interface{}) {
	fmt.Println("scan func")

}
