# AI Code Generation Edge Cases

## 1. AI-Generated Code with Syntax Errors
**Signal**: Codex for almost everything  
**Domain experiment**: Add `GenerateCode(pattern string) string` that returns AI-style code snippets with intentional syntax errors  
**Next failing test**: GenerateCode("for i in range(5): print(i)") = "for i in range(5): print(i)", want "for i in range(5): print(i)" (test error handling in downstream parsers)

## 2. Self-Hosting Interpreter with Variables
**