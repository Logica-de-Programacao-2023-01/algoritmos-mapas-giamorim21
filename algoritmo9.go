package main

import "fmt"

//Escreva uma função que gere a sequência de Fibonacci até um
//determinado número n e retorne um mapa onde as chaves são os
//índices da sequência e os valores são os números correspondentes.

func fibonacci(n int) map[int]int {
	Map := make(map[int]int)

	a, b := 0, 1
	for i := 0; i <= n; i++ {
		Map[i] = a
		a, b = b, a+b
	}
	return Map
}

func main() {
	n := 10
	Fibonacci := fibonacci(n)

	for i, valor := range Fibonacci {
		fmt.Printf("Fib(%d) = %d\n", i, valor)
	}
}
