# AI Code Generation Edge Cases → Domain Experiments

## 1. Codex for AI-Generated Code
**Signal**: AI-generated code patterns with edge cases  
**Domain experiment**: Add `GenerateCode(pattern string) string` that returns AI-style code snippets with syntax validation  
**Next failing test**: GenerateCode("def add(a, b): return a + b") = "def add(a, b): return a + b