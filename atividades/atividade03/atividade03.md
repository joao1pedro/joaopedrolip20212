# Nome: João Pedro

> Correção: 1,7
> 
> Faltou a árvore da segunda questão. Na terceira, faltou tratar o break. Mas o esforço foi legal, muito bom.

### Questão 01

 - 1. Derivação a esquerda. A = A + (B + C) 

```
<assign> => <id> = <expr>
         => A = <expr>
         => A = <expr> + <term>
         => A = <term> + <term>
         => A = <factor> + <term>
         => A = <id> + <term>
         => A = A + <term>
         => A = A + <factor>
         => A = A + (<expr>)
         => A = A + (<expr> + <term>)
         => A = A + (<term> + <term>)
         => A = A + (<factor> + <term>)
         => A = A + (<id> + <term>)
         => A = A + (B + <term>)
         => A = A + (B + <factor>)
         => A = A + (B + <id>)
         => A = A + (B + C)
```
<img src="./img/q1-a.jpg"
     alt="Q1-item1"/>

 - 2. Derivação a direita. A = B * (C * (A + B))

```
<assign> => <id> = <expr>
         => A = <expr>
         => A = <term>
         => A = <term> * <factor>
         => A = <id> * <factor>
         => A = B * <factor>
         => A = B * (<expr>)
         => A = B * (<term>)
         => A = B * (<term> * <factor>)
         => A = B * (<id> * <factor>)
         => A = B * (C * <factor>)
         => A = B * (C * (<expr>))
         => A = B * (C * (<expr> + <term>))
         => A = B * (C * (<term> + <term>))
         => A = B * (C * (<factor> + <term>))
         => A = B * (C * (<id> + <term>))
         => A = B * (C * (A + <term>))
         => A = B * (C * (A + <factor>))
         => A = B * (C * (A + <id>))
         => A = B * (C * (A + B))
```
<img src="./img/q1-b.jpg"
     alt="Q1-item2"/>

### Questão 02

Derivação a direita. (a 23 (m x y))

```
<lexp> => <lista> 
       => (<lexp-seq>)
       => (<lexp-seq> <lexp>)
       => (<lexp-seq> <lista>)
       => (<lexp-seq> (<lexp-seq>))
       => (<lexp-seq> (<lexp-seq> <lexp>))
       => (<lexp-seq> (<lexp-seq> <átomo>))
       => (<lexp-seq> (<lexp-seq> y))
       => (<lexp-seq> (<lexp-seq> <lexp> y))
       => (<lexp-seq> (<lexp-seq> <átomo> y))
       => (<lexp-seq> (<lexp-seq> x y))
       => (<lexp-seq> (<lexp> x y))
       => (<lexp-seq> (<átomo> x y))
       => (<lexp-seq> (m x y))
       => (<lexp-seq> <lexp> (m x y))
       => (<lexp-seq> <átomo> (m x y))
       => (<lexp-seq> 23 (m x y))
       => (<lexp> 23 (m x y))
       => (<átomo> 23 (m x y))
       => (a 23 (m x y))
```

### Questão 3

```
<switch_stmt> => switch ( identificador ) { <stmt> ; <stmt_list>}
<stmt_list>   => <stmt> <stmt_list>; | <stmt>
<stmt>        => <expr> |  <expr_list> | case <cst> : <stmt> | case <cst> : <stmt_list>
<expr_list>   => <expr> | ( <expr_list> ) | <stmt> | <expr> <expr>
<expr>        => <expr> + <expr> | <expr> - <expr> | <expr> * <expr> | <expr> / <expr> | <term>
<term>        => identificador | <cst> | identificador = <expr>; | <cst>; | identificador; identificador = <expr>; default: <stmt_list>
<cst>         => 1| 2| 3| ... n
```

- Switch Case avaliado:

```
switch (opcao) {
	case 1:
		soma = 2 + 5;
	case 2:
		soma = 3 + 2;
	default:
		soma = 0;
}
```

```
<switch_stmt> => switch ( id ) { <stmt>; <stmt_list>}
=> switch ( id ) { <expr_list> }
=> switch ( id ) { case <cst> : <stmt>; <stmt_list> }
=> switch ( id ) { case 1 : <stmt>; <stmt_list> }
=> switch ( id ) { case 1 : <expr_list>; <stmt_list> }
=> switch ( id ) { case 1 : <expr_list> <expr>; <stmt_list> }
=> switch ( id ) { case 1 : <expr> ; <stmt_list> }
=> switch ( id ) { case 1 : <term> ; <stmt_list> }
=> switch ( id ) { case 1 : id = <expr>; <stmt_list> }
=> switch ( id ) { case 1 : id = <expr> + <expr>; <stmt_list> }
=> switch ( id ) { case 1 : id = <term> + <expr>; <stmt_list> }
=> switch ( id ) { case 1 : id = <cst> + <expr>; <stmt_list> }
=> switch ( id ) { case 1 : id = 2 + 5; <stmt_list> }
=> switch ( id ) { case 1 : id = 2 + 5; <stmt> }
=> switch ( id ) { case 1 : id = 2 + 5; case <cst> : <stmt> }
=> switch ( id ) { case 1 : id = 2 + 5; case 2 : <stmt> }
=> switch ( id ) { case 1 : id = 2 + 5; case 2 : <expr_list> }
=> switch ( id ) { case 1 : id = 2 + 5; case 2 : <expr> }
=> switch ( id ) { case 1 : id = 2 + 5; case 2 : <term> }
=> switch ( id ) { case 1 : id = 2 + 5; case 2 : id = <expr>; default: <stmt_list>}
=> switch ( id ) { case 1 : id = 2 + 5; case 2 : id = <expr> + <expr>; default: <stmt_list>}
=> switch ( id ) { case 1 : id = 2 + 5; case 2 : id = <term> + <expr>; default: <stmt_list>}
=> switch ( id ) { case 1 : id = 2 + 5; case 2 : id = <cst> + <expr>; default: <stmt_list>}
=> switch ( id ) { case 1 : id = 2 + 5; case 2 : id = 2 + <expr>; default: <stmt_list>}
=> switch ( id ) { case 1 : id = 2 + 5; case 2 : id = 2 + <term>; default: <stmt_list>}
=> switch ( id ) { case 1 : id = 2 + 5; case 2 : id = 2 + <cst>; default: <stmt_list>}
=> switch ( id ) { case 1 : id = 2 + 5; case 2 : id = 2 + 3; default: <stmt_list>}
=> switch ( id ) { case 1 : id = 2 + 5; case 2 : id = 2 + 3; default: <stmt>}
=> switch ( id ) { case 1 : id = 2 + 5; case 2 : id = 2 + 3; default: <expr>}
=> switch ( id ) { case 1 : id = 2 + 5; case 2 : id = 2 + 3; default: <term>}
=> switch ( id ) { case 1 : id = 2 + 5; case 2 : id = 2 + 3; default: id = <expr>}
=> switch ( id ) { case 1 : id = 2 + 5; case 2 : id = 2 + 3; default: id = <term>}
=> switch ( id ) { case 1 : id = 2 + 5; case 2 : id = 2 + 3; default: id = <cst>}
=> switch ( id ) { case 1 : id = 2 + 5; case 2 : id = 2 + 3; default: id = 0;}
```
