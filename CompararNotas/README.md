# Desafio de comparar notas

Alice e Bob criaram cada um desafio. Um revisor avalia os dois desafios, atribuindo pontos em uma
escala de 1 a 100 para três categorias: clareza do problema , originalidade e dificuldade .

A classificação para o desafio de Alice é a lista ```a = (a[0], a[1], a[2])``` , e a classificação para o
desafio de Bob é a lista ```b = (b[0], b[1], b [2])``` .

A tarefa é encontrar seus pontos de comparação comparando a[0] com b[0] , a[1] com b[1] e a[2] com b[2] .

```Se a[i] > b[i]``` , então Alice ganha 1 ponto.  
```Se a[i] < b[i]``` , Bob recebe 1 ponto.  
```Se a[i] = b[i]``` , nenhuma pessoa recebe um ponto.  

Pontos de comparação é o total de pontos que uma pessoa ganhou.

Dados a e b , determine seus respectivos pontos de comparação.

Exemplo

```a = [1, 2, 3]```
```b = [3, 2, 1]```

Para os elementos ```*0*```, Bob recebe um ponto porque ```a[0] < b[0]```.
Para os elementos iguais ```a[1] e b[1]``` , nenhum ponto é ganho.
Finalmente, para os elementos ```2```, ```a[2] > b[2]``` então Alice recebe um ponto.
A matriz de retorno é ```[1, 1]``` com a pontuação de Alice primeiro e a de Bob em segundo.

Descrição da função

Complete a função deve Retornar

```int[2]``` : a pontuação de Alice está na primeira posição e a pontuação de Bob está na segunda.
Formato de entrada

A primeira linha contém 3 números inteiros separados por espaço, a[0] , a[1] e a[2] , os respectivos valores
no trio a .
A segunda linha contém 3 números inteiros separados por espaços, b[0] , b[1] e b[2] , os respectivos valores
no tripleto b .

Restrições
```
1 ≤ a[i] ≤ 100
1 ≤ b[i] ≤ 100
```
1 - Exemplo de Entrada
```
5 6 7
3 6 10
```
1 - Saída de amostra
```
1 1
```
2 - Exemplo de entrada
```
17 28 30
99 16 8
```
2 - Saída de amostra
```
2 1
```
