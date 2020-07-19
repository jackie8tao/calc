grammar calc;

expr: term prefixexp
    ;

prefixexp: '+' term prefixexp
    | '-' term prefixexp
    | /* epsilon */
    ;

term: factor prefixterm
    ;

prefixterm: '*' factor prefixterm
    | '/' factor prefixterm
    | /* epsilon */
    ;

factor: '(' expr ')'
    | INT
    ;

INT: [0-9]+;
