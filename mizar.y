translation_unit    ->  definition_or_statement
                    |   translation_unit definition_or_statement

definition_or_statement ->  function_definition
                        |   statement

function_definition ->  FUNC IDENTIFIER LP parameter_list RP block

parameter_list  ->  IDENTIFIER
                |   parameter_list COMMA IDENTIFIER

statement_list  ->  statement
                |   staement_list statement

statement -> expression SEMICOLON
        |   while_statement
        |   if_statement
        |   break_statement
        |   continue_statement
        |   return_statement
while_statement ->  WHILE LP expression RP block
if_statement -> IF LP expression RP block
            |   IF LP expression RP block ELSE block
            |   IF LP expression RP block elseif_list ELSE block

break_statement -> BREAK SEMICOLON

continue_statement -> CONTINUE SEMICOLON

return_statement -> RETURN expression SEMICOLON

elseif_list ->  elseif
            ->  elseif_list elseif

elseif  ->      ELSEIF LR expression RP block

expression ->   additive_expression
            |   IDENTIFIER ASSIGN additive_expression
            |   func_call_expression
additive_expression ->  multiplicative_expression
                    |   multiplicative_expression ADD additive_expression
                    |   multiplicative_expression SUB additive_expression
multiplicative_expression ->    primary_expression
                            |   primary_expression MUL multiplicative_expression
                            |   primary_expression DIV multiplicative_expression
primary_expression ->   STRING
                    |   NUMBER
                    |   IDENTIFIER
                    |   LP expression RP
                    |   func_call_expression

func_call_expression ->   IDENTIFIER Lp RP
                        | IDENTIFIER Lp argument_list RP

argument_list   ->  expression
                |   argument_list COMMA expression

block   ->  LC  statement_list RC
        |   LC RC


