package main

import "math"

//go test *.go
//go test pacote.go pacote_test.go

/*
Desafio Diferença entre Somas das Diagonais

Dada uma matriz quadrada, calcule a diferença absoluta entre as somas de suas diagonais.

Por exemplo, a matriz quadrada n = 3 é mostrado abaixo:

1 2 3
4 5 6
9 8 9

A diagonal da esquerda para a direita = {1, 5, 9}. A diagonal da direita para a esquerda = {3, 5, 9}. Sua diferença absoluta é |(1 + 5 + 9) - (3 + 5 + 9)| = |15 - 17| = 2

Descrição da função

A função recebe o seguinte parâmetro:

Uma matriz de inteiros

Retornar

Um valor inteiro com a diferença diagonal absoluta

Formato de saída

Retorna a diferença absoluta entre as somas das duas diagonais da matriz como um único inteiro.

Entrada de amostra

n = 3

11  2   4
4   5   6
10  8  -12

Saída de amostra

15

Explicação

A diagonal primária é:

11
   5
     -12

Soma na diagonal primária: 11 + 5 - 12 = 4

A diagonal secundária é:

     4
   5
10

Soma na diagonal secundária: 4 + 5 + 10 = 19 Diferença: |4 - 19| = 15

Observação: |x| é o valor absoluto de x


*/

func getPrimaryDiagonalSum(matrix [][]int) int {
	result := 0
	lenMatrix := len(matrix)

	// forma mais otimizada
	// for i, j := 0, 0; i < lenMatrix; i, j = i+1, j+1 {
	// 		result += matrix[i][j]
	// }

	for i := 0; i < lenMatrix; i++ {
		for j := 0; j < lenMatrix; j++ {
			if i == j {
				result += matrix[i][j]
			}
		}
	}

	return result
}

// 0,2
// 1,1
// 2,0
func getSecondaryDiagonalSum(matrix [][]int) int {
	result := 0
	lenMatrix := len(matrix)

	for i, j := 2, 0; i < lenMatrix || j > 0; i++,j-- {
		result += matrix[i][j]
	}
	return result
}

func GetDiagonalDiff(matrix [][]int) int {
	primaryDiagonal := getPrimaryDiagonalSum(matrix)
	secondaryDiagonal := getSecondaryDiagonalSum(matrix)
	return int(math.Abs(float64(primaryDiagonal) - float64(secondaryDiagonal)))
}
