// IntegraÃ§Ã£o numÃ©rica utilizando a regra dos trapÃ©zios.
// Adaptado de: http://linguagem-programacao-c.blogspot.com/2015/05/integracao-numerica-utilizando-regra_27.html
#include <stdio.h>
#include <math.h>

double f(double x);

int main() {
  // DeclaraÃ§Ã£o de variÃ¡veis.
    int i;
    int n = 0;      // NÃºmero de partiÃ§Ãµes.
    double a, b;    // Limites do intervalo de integraÃ§Ã£o.
    double sum = 0; // Valor do integral.
    double h;       // Altura.
    
    // Entrada de dados.
    printf("Este programa calcula o integral no intervalo [a,b]\n");
    printf("Introduza limite inferior a: \n");
    scanf("%lf", &a);
    
    printf("Introduza limite superior b (b>a): \n");
    scanf("%lf", &b);
    
    // ManipulaÃ§Ã£o da VariÃ¡vel "b"
    b = (5 * a)/2 - b;
    printf("Introduza nÃºmero de partiÃ§Ãµes do intervalo n (n > 1): \n");
    scanf("%d", &n);
    
    // Calcular dimensÃ£o de cada partiÃ§Ã£o.
    h = (b - a) / n ;
    printf("Valor de h: %f\n", h);
    
    // Ciclo de cÃ¡lculo.
    for (i = 0; i < n; i++) {
        sum = sum + (f(a) + f((a + h)))*h/2;
        a = a + h;
    }
    printf("\n");
    
    // Escrita do resultado.
    printf("O resultado da soma Ã©: %lf\n", sum);
} 
// end main()

// FunÃ§Ã£o a integrar.
double f(double x) {
    return  sqrt(1+ sin(x)*sin(x));
} // fim funcao.