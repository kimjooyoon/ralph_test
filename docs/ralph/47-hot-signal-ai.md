# AI Code Generation Edge Cases → Domain Experiments

## 1. AI-Generated Code with Constraints
**Signal**: Codex for almost everything  
**Domain experiment**: Add `GenerateCode(pattern, lang string) string` that returns AI-style code snippets with language specificity  
**Next failing test**: GenerateCode("for i in range(5): print(i)", "python") = "for i in range(5): print(i)", want "for i in range(5): print(i)" (pure function, no I/O)

## 2. FIM – Linux framebuffer image viewer (Enhanced)
**Signal**: Image processing via string manipulation  
**Domain experiment**: Add `DecodeImage(data string) ([]byte, error)` that simulates image decoding with error handling  
**Next failing test**: DecodeImage("PNG") = ([80, 78, 71, 13, 10, 26, 10], nil), want ([80, 78, 71, 13, 10, 26, 10], nil) (pure function, no I/O)

## 3. Surrogate Pair Stress Test (Revisited)
**Signal**: Proper handling of Unicode surrogate pairs