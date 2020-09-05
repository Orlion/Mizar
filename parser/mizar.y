translation_unit -> class_interface_declaration_list

class_interface_declaration_list ->     class_interface_declaration
                                |       class_interface_declaration_list class_interface_declaration

class_interface_declaration ->       class_declaration
                        |       interface_declaration
// 接口声明
interface_declaration ->        INTERFACE IDENTIFIER LC RC
                        |       INTERFACE IDENTIFIER LC interface_method_declaration_statement_list RC

// 接口内部方法声明
interface_method_declaration_statement_list ->  interface_method_declaration_statement
                                                | interface_method_declaration_statement_list interface_method_declaration_statement

interface_method_declaration_statement ->       type_var LP RP SEMICOLON
                                        |       type_var LR parameter_list RP SEMICOLON

// 类声明
class_declaration ->    CLASS IDENTIFIER LC RC
                |       ABSTRACT CLASS IDENTIFIER LC RC
                |       CLASS IDENTIFIER LC class_statement_list RC
                |       ABSTRACT CLASS IDENTIFIER LC class_statement_list RC
                |       CLASS IDENTIFIER extends_declaration LC RC
                |       ABSTRACT CLASS IDENTIFIER extends_declaration LC RC
                |       CLASS IDENTIFIER LC class_statement_list extends_declaration RC
                |       ABSTRACT CLASS IDENTIFIER extends_declaration LC class_statement_list RC
                |       CLASS IDENTIFIER implements_declaration LC RC
                |       ABSTRACT CLASS IDENTIFIER implements_declaration LC RC
                |       CLASS IDENTIFIER LC class_statement_list implements_declaration RC
                |       ABSTRACT CLASS IDENTIFIER implements_declaration LC class_statement_list RC
                |       CLASS IDENTIFIER extends_declaration implements_declaration LC RC
                |       ABSTRACT CLASS IDENTIFIER extends_declaration implements_declaration LC RC
                |       CLASS IDENTIFIER LC class_statement_list extends_declaration implements_declaration RC
                |       ABSTRACT CLASS IDENTIFIER extends_declaration implements_declaration LC class_statement_list RC

// extends声明
extends_declaration -> EXTENDS IDENTIFIER
                        | extends_declaration COMMA IDENTIFIER

implements_declaration -> IMPLEMENTS IDENTIFIER
                        | implements_declaration COMMA IDENTIFIER

class_statement_list -> class_statement
                        | class_statement_list class_statement

class_statement ->    method_definition
                |   property_definition

property_definition ->  member_modifier type_var SEMICOLON
                    |   member_modifier type_var ASSIGN expression_statement

method_definition ->    member_modifier type_var LP RP block
                        member_modifier type_var LP parameter_list RP block

member_modifier   ->      PUBLIC
                | PRIVATE
                | PROTECTED
                | ABSTRACT

parameter_list  ->  type_var
                |       parameter_list COMMA type_var

block   ->  LC  statement_list RC
        |   LC RC

statement_list  ->  statement
                |   statement statement_list

statement -> expression_statement
        |   var_declaration_statement
        |   var_assign_statement
        |   while_statement
        |   if_statement
        |   for_statement
        |   break_statement
        |   continue_statement
        |   return_statement
        | expression_statement

expression_statement -> expression SEMICOLON

while_statement ->  WHILE LP expression RP block
if_statement -> IF LP expression RP block
            |   IF LP expression RP block ELSE block

for_statement ->        FOR LR                  SEMICOLON               SEMICOLON               RP block //000
                |       FOR LR expression       SEMICOLON               SEMICOLON               RP block //100
                |       FOR LR expression       SEMICOLON expression    SEMICOLON               RP block //110
                |       FOR LR expression       SEMICOLON expression    SEMICOLON expression    RP block //111
                |       FOR LR expression       SEMICOLON               SEMICOLON expression    RP block //101
                |       FOR LR                  SEMICOLON               SEMICOLON expression    RP block //001
                |       FOR LR                  SEMICOLON expression    SEMICOLON expression    RP block //011
                |       FOR LR                  SEMICOLON expression    SEMICOLON               RP block //010

break_statement -> BREAK SEMICOLON
                |       BREAK expression_statement

continue_statement -> CONTINUE SEMICOLON
                |       CONTINUE expression_statement

return_statement -> RETURN SEMICOLON
                |       RETURN expression_statement
                        
var_declaration_statement -> type_var SEMICOLON

var_assign_statement -> type_var ASSIGN expression_statement // 变量声明并赋值
                        |       var_call_expression ASSIGN expression_statement // 给变量赋值

type_var -> IDENTIFIER IDENTIFIER // 变量声明 Int abc

expression ->           STRING_LITERAL
                    |   INT_LITERAL
                    |   DOUBLE_LITERAL
                    |   NULL
                    |   TRUE
                    |   FALSE
                    |   new_obj_expression
                    |   call_expression

call_expression -> var_call_expression
                -> method_call_expression

method_call_expression -> call_expression DOT method_call

var_call_expression -> call_expression DOT IDENTIFIER
                    |   THIS
                    |   IDENTIFIER

new_obj_expression -> NEW method_call  // new Class()

method_call ->  IDENTIFIER LP RP
            |   IDENTIFIER LP argument_list RP

argument_list   ->  expression
                | argument_list COMMA expression
