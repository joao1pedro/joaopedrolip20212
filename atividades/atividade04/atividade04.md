# Nome: João Pedro

> Correção: 1,5 

### Questão 1
<img src="./img/q1.jpeg"
     alt="Q1"/>

> Correção: 0,3
> 
> Onde estão as derivações?

### Questão 2
```
<program> -> begin <stmt_list> end 
<stmt_list> -> <stmt>
| <stmt> ; <stmt_list> 
<stmt> -> <var> = <expression>
<var> -> A | B | C
<expression> -> <var> {(+ | -) <var>}
```

> Correção: 0,2
> 
> Só fez \<var\> e não atualizou \<stmt_list\>?


### Questão 3
```
<assign> -> <id> = <expr> 
<id> -> A|B|C
<expr> -> <id> { (+ | - ) ( <expr> | (<expr>) ) }
```

> Correção: 0,5
> 
> OK!!!


### Questão 4
```
1. Regra sintática: <assign> -> <var> = <expr>
   Regra semântica: <expr>.expected_type <- <var>.actual_type
2. Regra sintática: <expr> -> <var>[2] + <var>[3]
   Regra semântica: <expr>.actual_type <-
               if (<var>[2].actual_type = int) and(<var>[3].actual_type = int)
                    then int
               else real
               end if
   Predicado: <expr>.actual_type == <expr>.expected_type | <expr>.actual_type != <expr>.expected_type
3. Regra sintática: <expr> -> <var>
   Regra semântica: <expr>.actual_type <- <var>.actual_type
   Predicado: <expr>.actual_type == <expr>.expected_type | <expr>.actual_type != <expr>.expected_type
4. Regra sintática: <var> -> A | B | C
   Regra semântica: <var>.actual_type <- look-up(<var>.string)
```

> Correção: 0,3
> 
> O predicado ficou estranho. O correto seria apenas retirá-lo.
