# calc
`calc` is a simple calculator implemented in compiler theory.

## grammar

raw expression:

```antlrv4
grammar calc;
expr: expr '+' term
    | expr '-' term
    | term;
term: term '*' factor
    | term '/' factor
    | factor;
factor: '(' expr ')'
    | INT;
INT: [0-9]+;
```

formatted expression:

```antlrv4
grammar calc;
expr: term prefixexp;
prefixexp: '+' term prefixexp
    | '-' term prefixexp
    | /* epsilon */;
term: factor prefixterm;
prefixterm: '*' factor prefixterm
    | '/' factor prefixterm
    | /* epsilon */;
factor: '(' expr ')'
    | INT;
INT: [0-9]+;
```

## author

- name: Jackie
- email: jackie8tao@outlook.com
