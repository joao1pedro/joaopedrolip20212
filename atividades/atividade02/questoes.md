# Nome: João Pedro

 ### 1. Os três principais métodos de implementação de linguagens de programação são:
 - **Compilação:** Nesta abordagem, os programas são transformados em intruções de linguagem de máquina (linguagem que o computador consegue reconhecer) a partir da compilação. Exemplos de linguagens nesta categoria são: C, C++ e COBOL.
 - **Interpretação:** Neste caso, os programas não são traduzidos como ocorre na compilação, os programas são interpretados por um programa especial que realiza essa função de interpretar o código fonte do programa. Exemplos de linguagens que utilizam esta abordagem são: Python, PHP, Javascript.
 - **Sistemas de implementação hibrida:** Nesta abordagem, há um meio termo entre a compilação e interpretação. Os programas são traduzidos para uma linguagem intermediária de interpretação. Um exemplo de linguagem que se encaixa nesta categoria é o Java.

 ### 2. O comentário em múltiplas linhas tem a vantagem de ser muito utilizado para documentar códigos fontes de aplicações uma vez que podem explicar de uma só vez o que a função recebe de parâmetro, o que retorna como resultado e o que a função faz. Como desvantagem, este tipo de comentário pode poluir muito o código se o programador utilizar este recurso demasiadamente. Já o segunddo caso onde tem-se um comentário por linha é muito utilizado para comentários breves, como por exemplo um lembrete sobre o que um laço for está fazendo, ou porque determinada variavel é incrementada. Como desvantagem, as vezes são utilizados muitos comentários de linha única no lugar de um comentário de várias linhas, algo que deve ser levado em consideração pelo programador.

 ### 3. O FORTRAN e ALGOL foram feitos para utilizar aplicações ciêntificas onde utilizam-se basicamente calculos vetoriais, matriciais e aritméticos de ponto flutuante (maior precisão). Sendo assim, o quanto mais simples os identificadores, melhor é para os cientistas/desenvolvedores que trabalham com estas linguagens. O COBOL foi criado pensando em aplicações empresáriais, sendo assim, são utilizados para produzir relatórios, trabalhar com caracteres numéricos e strings, e cálculos com números decimais. Sendo assim, o uso de identificadores maiores permite maior que as estruturas codificadas sejam identificadas com maior facilidade, é melhor para a identificação de métodos e/ou funções mais complexas.

 ### 4. Baseado em contagem:
 ```
for(int i = 0; i < linhas; i++) {
    for (int j = 0; j < colunas; j++)
    {
        for (int k = 0; k < colunas; k++)
        {
            aux = aux + (matriz1[i][k] * matriz2[k][j]);
        }
        matrizMult[i][j] = aux;
        aux = 0;
    }
}
 ```
 Nesta abordagem as matrizes são representadas por vetores de duas dimensões.

### com while:
```
while (i < linhas) {
    i++;
    while (j < colunas)
    {
        j++;
        while (k < colunas)
        {
            k++;
            mAux = mAux + (matriz1[i][k] * matriz2[k][j]);
        }
        matrizMult[i][j] = mAux;
        mAux = 0;
    }
}
```
No caso da implementação com while, os indices tem que ser incrementados manualmente, algo que não ocorre no for, onde isso já ocorre de maneira automática a cada iteração.