package main

import "fmt"

func main() {
	type NoKTP string

	var ktpilham NoKTP = "1234567890"

	var contoh = "9999999999"

	fmt.Println(ktpilham)
	var noktp = NoKTP(contoh)
	fmt.Println(noktp)

}
