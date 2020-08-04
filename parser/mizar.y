translation_unit -> class_declaration_list

class_declaration_list ->  class_declaration
                  |  class_declaration_list class_declaration

method_definition ->  (ABSTRACT) (PUBLIC|PRIVATE|PROTECTED) IDENTIFIER IDENTIFIER LP parameter_list RP block

parameter_list  ->  IDENTIFIER IDENTIFIER
                |   IDENTIFIER IDENTIFIER COMMA parameter_list

// 接口声明
interface_declaration ->  INTERFACE IDENTIFIER LC RC
                        | INTERFACE IDENTIFIER LC interface_function_declaration_statement_list RC

// 接口内部方法声明
interface_method_declaration_statement_list ->  interface_function_declaration_statement
                                                | interface_function_declaration_statement_list interface_function_declaration_statement

interface_method_declaration_statement -> IDENTIFIER IDENTIFIER LR RP SEMICOLON
                                | void IDENTIFIER LR RP SEMICOLON
                                | IDENTIFIER IDENTIFIER LR parameter_list RP SEMICOLON
                                | void IDENTIFIER LR parameter_list RP SEMICOLON

// extends声明
extends_declaration -> EXTENDS IDENTIFIER
                        | extends_declaration IDENTIFIER

implements_declaration -> IMPLEMENTS IDENTIFIER
                        | implements_declaration IDENTIFIER

// 类声明
class_declaration -> ABSTRACT CLASS IDENTIFIER LC  RC
                |    ABSTRACT CLASS IDENTIFIER LC class_statement_list RC
                | CLASS IDENTIFIER LC  RC
                | CLASS IDENTIFIER LC class_statement_list RC

class_statement_list -> class_statement
                        | class_statement_list class_statement

class_statement ->    method_definition
                | PUBLIC var_declaration SEMICOLON
                | PRIVATE var_declaration SEMICOLON
                | PROTECTED var_declaration SEMICOLON
                | CONST var_declaration SEMICOLON

// 变量声明
value_declaration ->      IDENTIFIER IDENTIFIER

statement_list  ->  statement
                |   statement staement_list

statement -> expression SEMICOLON
        |   IDENTIFIER ASSIGN expression SEMICOLON
        |   IDENTIFIER DOT IDENTIFIER ASSIGN expression SEMICOLON
        |   var_declaration ASSIGN expression SEMICOLON // 声明并赋值
        |   while_statement
        |   if_statement
        |   for_statement
        |   break_statement
        |   continue_statement
        |   return_statement
        |   var_declaration_statement

value_declaration_statement -> value_declaration SEMICOLON

while_statement ->  WHILE LP expression RP block
if_statement -> IF LP expression RP block
            |   IF LP expression RP block ELSE block

for_statement -> FOR LR expression_opt SEMICOLON expression_opt SEMICOLON expression_opt block

break_statement -> BREAK expression_opt SEMICOLON

continue_statement -> CONTINUE expression_opt SEMICOLON

return_statement -> RETURN expression_opt SEMICOLON
                        
expression_opt -> // NULL
                | expression

method_call_expression ->   primary_expression DOT IDENTIFIER Lp RP
                        | primary_expression DOT IDENTIFIER LP argument_list RP
                        | IDENTIFIER DOUBLE_COLON IDENTIFIER Lp RP
                        | IDENTIFIER DOUBLE_COLON IDENTIFIER Lp argument_list RP

argument_list   ->  primary_expression
                |   primary_expression COMMA argument_list

block   ->  LC  statement_list RC
        |   LC RC


new_obj_expression -> NEW IDENTIFIER LP argument_list RP  // new Class(argument_list)
                    |   NEW IDENTIFIER LP RP  // new Class()

expression ->           STRING_LITERAL
                    |   INT_LITERAL
                    |   DOUBLE_LITERAL
                    |   NULL
                    |   TRUE
                    |   FALSE
                    |   IDENTIFIER
                    |   IDENTIFIER DOT IDENTIFIER // a.b访问对象属性
                    |   IDENTIFIER DOUBLE_COLON IDENTIFIER // a::b 访问类常量
                    |   new_obj_expression
                    |   method_call_expression
                    