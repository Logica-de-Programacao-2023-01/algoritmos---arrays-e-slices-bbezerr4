package main

import "fmt"

//Crie um Array de floats com 4 elementos e calcule o produto dos valores armazenados no Array.

func main() {
	numeros := [4]float64{4.4, 6.7, 7.6, 8.6}
	produto := 1.0

	for _, num := range numeros {
		produto = num * produto

	}
	fmt.Print(produto)
}
