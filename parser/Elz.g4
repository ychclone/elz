grammar Elz;

options {
    language = Go;
}

WS: [ \t\r\n]+ -> channel(HIDDEN);
COMMENT: '//' .*? '\n' -> channel(HIDDEN);

ID : StartLetter Letter*;
fragment
StartLetter: [a-zA-Z_]
    | '\u00C0'..'\u00D6'
    | '\u00D8'..'\u00F6'
    | '\u00F8'..'\u02FF'
    | '\u0370'..'\u037D'
    | '\u037F'..'\u1FFF'
    | '\u200C'..'\u200D'
    | '\u2070'..'\u218F'
    | '\u2C00'..'\u2FEF'
    | '\u3001'..'\uD7FF'
    | '\uF900'..'\uFDCF'
    | '\uFDF0'..'\uFFFD'
    ;
fragment
Letter: StartLetter
    | [0-9]
    | '\u00B7'
    | '\u0300'..'\u036F'
    | '\u203F'..'\u2040'
    ;

NUM: StartDigit Digit*;
fragment
StartDigit: [0-9.];
fragment
Digit: [0-9];

STRING: '"' .*? '"';

prog: topStatList?;

topStatList: topStat+;

topStat: fnDefine // fn foo( params ) { stat... }
    | varDefine   // let (mut) var: type = expr
    | typeDefine  // type newType ( prop... )
    | implBlock   // impl type { method... }
    | traitDefine // trait DB { method... }
    | importStat  // import ( Module... )
    ;

importMod: ID ('::' ID)*;
importStat: 'import' '(' importMod ')';

statList: stat+;
stat: varDefine
    | loopStat // loop { stats }
    | returnStat
    | assign
    | exprStat
    ;

returnStat:
    'return' expr
    ;

loopStat:
    'loop' '{'
        statList?
    '}'
    ;

exprStat: matchRule
    | fnCall
    ;

// match i {
//   10 => { break },
//   _ => { i = i + 1 }
// }
matchRule:
    'match' expr '{'
        expr '=>' stat
        (',' expr '=>' stat)*
        ','?
    '}'
    ;

// var = 1
assign:
    ID '=' expr
    ;

// 1, 2, 3, 4, 5
exprList: expr (',' expr)*;
// add(1, 2)
fnCall:
    ID '(' exprList? ')'
    ;

// mean a type, but type already be use by Go, so need an alternative name
typePass : ID;
typeList: typePass (',' typePass)*;

// @op
annotation: '@' ID ('(' expr ')')? ;

methodList: method+;
method:
    exportor? ID '(' paramList? ')' ('->' typePass)? '{'
        statList?
    '}'
    ;
implBlock:
    'impl' ID (':' typeList)? '{'
        methodList?
    '}'
    ;

// exportor can use on detect variable scope
// Because local scope can't export, it may safe.
// And create another rule may to complex.
exportor: '+';
define: exportor? ID (':' typePass)? '=' expr;
varDefine:
    'let' mut='mut'? define (',' define)*
    ;
paramList: param (',' param)*;
param: ID (':' typePass)?;
fnDefine:
    // because fn also handle operator, so if we use exportor after keyword fn will cause we hard to divide ++ && + +
    exportor? 'fn' ID '(' paramList? ')' ('->' typePass)? '{'
        statList?
    '}'
    ;
attrList: attr+;
attr: exportor ID ':' typePass;
typeDefine:
    'type' exportor? ID '(' attrList ')'
    ;
tmethodList: tmethod+;
tmethod: exportor? ID '(' typeList? ')' ('->' typePass)?;
traitDefine:
    'trait' exportor ID '{'
        tmethodList?
    '}'
    ;

// Explain for expr, because Antlr support the operation precedence by declared order
// So we don't have to consider that
expr: expr op='^' expr       # Pow // TODO: We had not support translate it.
    | expr op=('*'|'/') expr # MulOrDiv // operation prec
    | expr op=('+'|'-') expr # AddOrSub
    | expr op='!=' expr      # NotEq // TODO: Waiting for implement
    | expr op='==' expr      # Eq // TODO: Waiting for implement
    | expr '?' expr ':' expr # ThreeOpCmp // TODO: We had not support translate it.
    | '(' expr ')'           # SubExpr // TODO: Waiting for implement
    | exprStat               # StatExpr // Important, exprStat have match & functionCall yet!
    | NUM                    # Num
    | ID                     # Id
    | STRING                 # Str
    ;
