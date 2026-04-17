# HN: AI Code Generation → Domain Experiments

## 1. Claude Design AI Code Patterns
**Signal**: AI-generated code patterns from Claude Design
**Domain experiment**: Test `GenerateCode("for i in range(5): print(i)")` returns Python snippet
**Next failing test**: GenerateCode("for i in range(5): print(i)") = "for i in range(5): print(i)", want "for i in range(5): print(i)" (pure function, no I/O)

## 2. Self-Hosting Interpreter
**Signal**: Python interpreter in Python
**Domain experiment**: Test `Eval("1+1")` returns "2"