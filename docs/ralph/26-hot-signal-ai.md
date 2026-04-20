# AI-Generated Code Patterns → Domain Experiments

## 1. AI-Generated Code Snippets
**Signal**: Codex for almost everything  
**Domain experiment**: Add `GenerateCode(pattern string) string` that returns AI-style code snippets  
**Next failing test**: GenerateCode("for i in range(5): print(i)") = "for i in range(5): print(i)", want "for i in range(5): print(i)" (pure function, no I/O)

## 2. Self-Hosting Interpreter
**Signal**: Python interpreter in Python  
**Domain experiment**: Add `Eval(code string) string` that returns evaluated code (mocked)  
**Next failing test**: Eval("1+1") = "2", want "2" (pure function, no I/O)

## 3. Image Decoding Mock
**Signal**: FIM – Linux framebuffer viewer  
**Domain experiment**: Add `DecodeImage(data string) []byte` that simulates image decoding  
**Next failing test**: DecodeImage("PNG") = [80, 78, 71, 13, 10, 26, 10], want [