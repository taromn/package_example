package main

import (
	"fmt"

	"github.com/learning-go-book/package_example/formatter"  // path to package
	"github.com/learning-go-book/package_example/math"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)  // print -> package print in ./formatter
	fmt.Println(output)
}
