package main

import "fmt"

//Escreva uma função que receba dois mapas e retorne um novo
//mapa contendo todos os elementos dos mapas de entrada. Em caso
//de chaves duplicadas, o valor do segundo mapa deve prevalecer.

func Maps(map1, map2 map[string]int) map[string]int {
	Map := make(map[string]int)

	for k, valor := range map1 {
		Map[k] = valor
	}
	for k, valor := range map2 {
		Map[k] = valor
	}
	return Map
}

func main() {
	map1 := map[string]int{
		"apple":  1,
		"banana": 2,
	}
	map2 := map[string]int{
		"banana": 3,
		"orange": 4,
	}
	m := Maps(map1, map2)
	fmt.Println(m)
}
