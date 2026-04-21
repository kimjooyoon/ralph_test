# AI-Generated Code Patterns → Domain Experiments

## 6. AI-Generated Code Snippets (Revisited)
**Signal**: Codex for almost everything  
**Domain experiment**: Add `GenerateCode(pattern string) string` that returns AI-style code snippets  
**Next failing test**: GenerateCode("for i in range(5): print(i)") = "for i in range(5): print(i)", want "for i in range(5): print(i)" (pure function, no I/O)

## 7. Self-Hosting Interpreter (Revisited)
**Signal**: Python interpreter in Python  
**Domain experiment**: Add `Eval(code string) string` that returns evaluated code (mocked)  
**Next failing test**: Eval("1+1") = "2", want "2" (pure function, no I/O)

## 8. Regex-Like Pattern Matching (Revisited)
**Signal**: Lightweight regex alternatives  
**Domain experiment**: Add `Match(pattern, s string) bool` for basic regex-like checks  
**Next failing test**: Match("^[a-z]+$", "abc123") = false, want false (test