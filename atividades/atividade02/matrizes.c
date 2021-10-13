#define linhas 2
#define colunas 2

#include <stdio.h>

int main()
{
    int matriz1[linhas][colunas];
    int matriz2[linhas][colunas];
    int matrizMult[linhas][colunas];

    printf("\nPreencha a matriz 1: \n");
    for (int i = 0; i < linhas; i++)
    {
        for (int j = 0; j < colunas; j++)
        {
            printf("Digite o número da posição [%d, %d]: ", i, j);
            scanf("%d", &matriz1[i][j]);
        }
    }

    printf("\nPreencha a matriz 2: \n");
    for (int i = 0; i < linhas; i++)
    {
        for (int j = 0; j < colunas; j++)
        {
            printf("Digite o número da posição [%d, %d]: ", i, j);
            scanf("%d", &matriz2[i][j]);
        }
    }

    printf("\nMatriz 1: \n");
    for (int i = 0; i < linhas; i++)
    {
        for (int j = 0; j < colunas; j++)
        {
            printf(" %d ", matriz1[i][j]);
        }
        printf("\n");
    }

    printf("\nMatriz 2: \n");
    for (int i = 0; i < linhas; i++)
    {
        for (int j = 0; j < colunas; j++)
        {
            printf(" %d ", matriz2[i][j]);
        }
        printf("\n");
    }

    int aux = 0;
    printf("\nMultiplicação de matriz1 por matriz2:\n");
    for (int i = 0; i < linhas; i++)
    {
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

    printf("\nMatriz Multiplicação com laço FOR: \n");
    for (int i = 0; i < linhas; i++)
    {
        for (int j = 0; j < colunas; j++)
        {
            printf(" %d ", matrizMult[i][j]);
        }
        printf("\n");
    }

    int mAux, i, j, k = 0;
    printf("\nMultiplicação de matriz1 por matriz2:\n");
    while (i < linhas)
    {
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

    printf("\nMatriz Multiplicação com laço WHILE: \n");
    for (int i = 0; i < linhas; i++)
    {
        for (int j = 0; j < colunas; j++)
        {
            printf(" %d ", matrizMult[i][j]);
        }
        printf("\n");
    }

    return 0;
}
