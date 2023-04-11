package main

import "fmt"

//Crie um Array de inteiros com 3 elementos e imprima a soma dos valores armazenados no Array.

func main() {
	numeros := [3]int{1, 6, 7}

	soma := 0
	for _, num := range numeros {
		soma += num
	}

	fmt.Println(soma)
}
