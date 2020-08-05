translation_unit -> class_interface_declaration_list

class_interface_declaration_list ->     class_interface_declaration
                                |       class_interface_declaration_list class_interface_declaration

class_interface_declaration ->       class_declaration
                        |       interface_declaration
// 接口声明
interface_declaration ->  INTERFACE IDENTIFIER LC interface_function_declaration_statement_list RC

// 接口内部方法声明
interface_method_declaration_statement_list ->  interface_function_declaration_statement
                                                | interface_function_declaration_statement_list interface_function_declaration_statement

interface_method_declaration_statement -> return_val_type IDENTIFIER LR parameter_list RP SEMICOLON

// extends声明
extends_declaration -> EXTENDS IDENTIFIER
                        | extends_declaration IDENTIFIER

implements_declaration -> IMPLEMENTS IDENTIFIER
                        | implements_declaration IDENTIFIER

// 类声明
class_declaration ->    CLASS IDENTIFIER LC RC
                        ABSTRACT CLASS IDENTIFIER LC class_statement_list RC

class_statement_list -> class_statement
                        | class_statement_list class_statement

class_statement ->    method_definition
                | var_modifier var_declaration SEMICOLON
                | var_modifier var_declaration ASSIGN expression SEMICOLON
        
var_modifier   ->      PUBLIC
                | PRIVATE
                | PROTECTED

method_definition ->    method_modifier return_val_type IDENTIFIER LP RP block
                        method_modifier return_val_type IDENTIFIER LP parameter_list RP block

return_val_type -> void
                | IDENTIFIER

parameter_list  ->  IDENTIFIER IDENTIFIER
                |       IDENTIFIER IDENTIFIER COMMA parameter_list

method_modifier -> PUBLIC
                | PRIVATE
                | PROTECTED
                | ABSTRACT

block   ->  LC  statement_list RC
        |   LC RC

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
                        | 

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
                |       BREAK expression SEMICOLON

continue_statement -> CONTINUE SEMICOLON
                |       CONTINUE expression SEMICOLON

return_statement -> RETURN SEMICOLON
                |       RETURN expression SEMICOLON
                        
var_declaration_statement -> var_declaration SEMICOLON

var_assign_statement -> var_declaration ASSIGN expression SEMICOLON // 变量声明并赋值
                        |       var_call_expression ASSIGN expression SEMICOLON // 给变量赋值

var_declaration -> IDENTIFIER IDENTIFIER // 变量声明 Int abc

expression ->           STRING_LITERAL
                    |   INT_LITERAL
                    |   DOUBLE_LITERAL
                    |   NULL
                    |   TRUE
                    |   FALSE
                    |   IDENTIFIER
                    |   new_obj_expression
                    |   call_expression

call_expression -> var_call_expression
                -> method_call_expression


method_call_expression -> call_expression DOT IDENTIFIER Lp RP
                |       call_expression DOT IDENTIFIER Lp argument_list RP

var_call_expression -> IDENTIFIER
        |       THIS
        |       call_expression DOT IDENTIFIER

new_obj_expression -> NEW IDENTIFIER LP RP  // new Class()
                |       NEW IDENTIFIER LP argument_list RP  // new Class(argument_list)

argument_list   ->  expression
                | argument_list COMMA expression