 # Questão 01

Faça um programa que leia um número inteiro e o imprima.

<details>
  <summary>Resposta</summary>

```c
#include <stdio.h>

int main() {
    int x;
    //scanf exige "%(tipo da variável)", &(variável)
    //%d é usado para inteiros
    scanf("%d", &x);
    //no printf a cada %(tipo da variável) encontrado 
    //dentro das aspas, uma varível irá substituí-lo 
    //na ordem que é dada.
    printf("%d\n", x);
    return 0;
}
```

</details>

## Casos de Teste

```
>>>>>>>> INSERT
10
======== EXPECT
10
<<<<<<<< FINISH

>>>>>>>> INSERT
123
======== EXPECT
123
<<<<<<<< FINISH

>>>>>>>> INSERT
-45
======== EXPECT
-45
<<<<<<<< FINISH

>>>>>>>> INSERT
0
======== EXPECT
0
<<<<<<<< FINISH

>>>>>>>> INSERT
999999
======== EXPECT
999999
<<<<<<<< FINISH
```

[Código](../../.tko/cache/4753b204da0db175a0d57bf6679cc041a0f233d2/base/inteiro/./C/draft.c)
