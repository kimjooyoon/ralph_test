# HN: AI Code Generation → Domain Experiments

## 1. Claude Design AI Code Patterns
**Signal**: AI-generated code patterns from Claude Design  
**Domain experiment**: Add `GenerateCode(pattern string) string` that returns AI-style code snippets  
**Next failing test**: GenerateCode("for i in range(5): print(i)") = "for i in range(5): print(i)", want "for i in range(5): print(i)" (pure function, no I/O)

## 2. Self-Hosting Interpreter  
**Signal**: Python interpreter in Python  
**Domain experiment**: Add `Eval(code string) string` that returns evaluated code (mocked)  
**