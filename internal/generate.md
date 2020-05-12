# Generate ANTLR4 Code

ANTLR 4.8 is required for generating the source code. The following
command must be invoked to generate the source code:

`antlr -Dlanguage=Go -o parser -package parser -no-listener -visitor FHIRPath.g4`

