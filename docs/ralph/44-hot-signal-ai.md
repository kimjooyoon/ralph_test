# AI-Generated Code Patterns → Domain Experiments

## Signal: AI-generated code patterns
**Domain experiment**: Add `GenerateCode(pattern string) string` that returns AI-style code snippets  
**Next failing test**: GenerateCode("for i in range(5): print(i)") = "for i in range(5): print(i)", want "for i in range(5): print(i)" (pure function, no I/O)

## Signal: Self-hosted language interpreters
**Domain experiment**: Add `Eval(code string) string` that returns evaluated code (mocked)  
**Next failing test**: Eval("1+1") = "2", want "2" (pure function, no I/O)

## Signal: Image processing via string manipulation
**Domain experiment**: Add `DecodeImage(data string) []byte` that simulates image decoding  
**Next failing test**: DecodeImage("PNG") = [