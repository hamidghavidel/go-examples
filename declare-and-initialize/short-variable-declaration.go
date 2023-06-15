package main

import "fmt"

func main() {
	aa := 10
	bb := "hello"
	cc := 3.14
	dd := true

	fmt.Printf("aa:= 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb:= \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)
}
