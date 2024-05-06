package main

import "fmt"

func main() {

	// operasi boolean
	fmt.Println("true && true :", true && true)
	fmt.Println("true && false :", true && false)
	fmt.Println("true || true :", true || true)
	fmt.Println("true || false :", true || false)
	fmt.Println("!true :", !true)
	println("+++++++++++++++++++++++++++++++++++++++++++++")
	// operasi boolean sederhana
	nilaiuts := 80
	nilaiuas := 90

	var lutspassing bool = nilaiuts >= 75
	var luaspassing bool = nilaiuas >= 80

	var lulus bool = lutspassing && luaspassing

	fmt.Println("Lulus :", lulus)

}
