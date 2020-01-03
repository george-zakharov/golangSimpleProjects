// Go имеет различные типы данных, такие как строки,
// целые числа, числа с плавающей запятой,
// булев тип, и так далее.
// Вот несколько основных примеров.

package main

import "fmt"

func main() {

	// Строки можно объединять с помощью `+`.
	fmt.Println("go" + "lang")

	// Целые числа и числа с плавающей запятой.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Булев тип, как вы и ожидаете, с логическими операторами true или false.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}