package main

import "fmt"

func main() {
	a := 10
	b := 5
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)

	// augmented assignment
	a += 5
	fmt.Println("a += 5 :", a)
	a -= 5
	fmt.Println("a -= 5 :", a)
	a *= 5
	fmt.Println("a *= 5 :", a)
	a /= 5
	fmt.Println("a /= 5 :", a)
	a %= 5
	fmt.Println("a %= 5 :", a)

	// unary operator
	a++
	fmt.Println("a++ :", a)
	a--
	fmt.Println("a-- :", a)

}
