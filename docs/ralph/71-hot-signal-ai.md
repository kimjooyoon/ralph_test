# AI-Generated Code Validation → Domain Experiments

## 1. Codex Output Fidelity
**Signal**: AI-generated code patterns  
**Domain experiment**: Add `ValidateCode(pattern, output string) bool` to check AI-generated code against expected patterns  
**Next failing test**: ValidateCode("for i in range(5): print(i)", "for i in range(5): print(i)") = true, want false (test for strict pattern matching)

## 2. Code Generation with Context
**Signal**: Context-aware AI code generation  
**Domain experiment**: Add `GenerateCodeWithContext(prompt, context string) string` that incorporates contextual hints  
**Next failing test**: GenerateCodeWithContext("reverse string", "s = 'abc'") = "s = 'cba'", want "s = 'cba'" (pure function,