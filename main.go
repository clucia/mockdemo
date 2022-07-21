package main

import "github.com/clucia/mockdemo/scanner"

func main() {
	mys := &scanner.MyScanner{}

	var x, y int

	mys.Scan(&x, &y)

}
