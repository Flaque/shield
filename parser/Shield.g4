grammar Shield ;

shield : math ;

equals : math | math IS math ;
math : term | math '+' term | math '-' term ;
term : LABEL | INT | STRING ;

NEWLINE : [\r\n]+ ;
WHITESPACE : ' ' -> skip ;

STRING : '"' ALPHA '"' ;
LABEL : ALPHA ;
INT : [0-9]+ ; 
ALPHA : [a-zA-Z]+ ;
IS : 'is' ;