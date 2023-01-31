# Desafio Diferença entre Somas das Diagonais

Dada uma matriz quadrada, calcule a diferença absoluta entre as somas de suas diagonais.

Por exemplo, a matriz quadrada n = 3 é mostrado abaixo:
```
1 2 3
4 5 6
9 8 9
```
A diagonal da esquerda para a ```direita = {1, 5, 9}```. A diagonal da direita para a ```esquerda = {3, 5, 9}```.
Sua diferença absoluta é
```|(1 + 5 + 9) - (3 + 5 + 9)| = |15 - 17| = 2```

Descrição da função

A função recebe o seguinte parâmetro:

Uma matriz de inteiros

Retornar

Um valor inteiro com a diferença diagonal absoluta

Formato de saída

Retorna a diferença absoluta entre as somas das duas diagonais da matriz como um único inteiro.

Entrada de amostra
```
n = 3

11  2   4
4   5   6
10  8  -12
```

Saída de amostra
```
15
```
Explicação

A diagonal primária é:
```
11
   5
     -12
```
Soma na diagonal primária: ```11 + 5 - 12 = 4```

A diagonal secundária é:
```
     4
   5
10
```
Soma na diagonal secundária: ```4 + 5 + 10 = 19```
Diferença: ```|4 - 19| = 15```

Observação: ```|x|``` é o valor absoluto de ```x```
