php
===

Parser for PHP written in Go


## Project Status

This project is under heavy development, though some pieces are more or less stable. Listed here are components that in progress or are ideas for future development

Feature                       |Status
------------------------------|------
Lexer and Parser              | mostly complete. there are probably a few gaps still
Scoping                       | complete for simple cases. probably some gaps still, most notably that conditional definitions are treated as if they are always defined
Code search and symbol lookup | basic idea implemented, many many details missing
Code formatting               | basic idea implemented, formatting needs to narrow down to PSR-2
Transpilation to Go           | basic idea implemented, need follow through with more node types
Type inferencing              | not begun
Dead code analysis            | basic idea implemented, but only for some types of code. Also, this suffers from the same caveats as scoping

## Project Components

Directory                     |Description
------------------------------|------
php| the core parser
php/ast| (abstract syntax tree) describes the nodes in PHP as parsed by the parser
php/ast/printer| prints an ast back to source code
php/cmd| a tool used to debug the parser
php/lexer| reads a stream of tokens from source code
php/passes| tools and packages related to modifying or analyzing PHP code (heavily a work in progress)
php/passes/togo| transpiler
php/passes/deadcode| dead code analyzer
php/query| tools and packages related to analyzing and finding things in PHP code (heavily a work in progress)
php/testdata| simple examples of PHP that must parse with no errors for tests to pass
php/token| describes the tokens read by the lexer

original fork: https://github.com/stephens2424/php
