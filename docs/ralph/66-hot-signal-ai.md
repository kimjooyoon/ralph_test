# AI Code Generation → Domain Experiments

## 1. AI-Generated Code Patterns
**Signal**: AI-generated code patterns  
**Domain experiment**: Add `GenerateCode(pattern string) string` that returns AI-style code snippets  
**Next failing test**: GenerateCode("for i in range(5): print(i)") = "for i in range(5): print(i)", want "for i in range(5): print(i)" (pure function, no I/O)

## 2. Self-Hosted Interpreter
**Signal**: Self-hosted language interpreters  
**Domain experiment**: Add `Eval(code string) string` that returns evaluated code (mocked)  
**Next failing test**: Eval("1+1") = "2", want "2" (pure function, no I/O)

## 3. Regex-Like Matching
**Signal**: Lightweight regex alternatives  
**Domain experiment**: Add `Match(pattern, s string) bool` for basic regex-like checks  
**Next failing test**: Match("^[a-z]+$", "abc123") = false, want false (test for simple regex syntax)
