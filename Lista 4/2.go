package main

import "fmt"

func main() {
	vetor1 := make([]int, 10)
	vetor2 := make([]int, 5)
	fmt.Print("Primeiro vetor:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&vetor1[i])
	}
	soma := 0
	fmt.Print("Segundo vetor:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&vetor2[i])
		soma = soma + vetor2[i]
	}
	pares := 0
	impares := 0
	for i := 0; i < 10; i++ {
		if vetor1[i]%2 == 0 {
			pares = pares + 1
		} else {
			impares = impares + 1
		}
	}
	vetor_pares := make([]int, pares)
	vetor_impares := make([]int, impares)
	j := 0
	k := 0
	for i := 0; i < 10; i++ {
		if vetor1[i]%2 == 0 {
			vetor_pares[j] = vetor1[i] + soma
			j = j + 1
		} else {
			vetor_impares[k] = vetor1[i] + soma
			k = k + 1
		}
	}
	fmt.Println("Primeiro vetor resultante:", vetor_pares)
	fmt.Print("Segundo vetor resultante:", vetor_impares)
}
