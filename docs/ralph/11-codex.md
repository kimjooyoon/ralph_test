# Codex for almost everything → Domain Experiments

## 2. Codex for almost everything
**Signal**: AI-generated code patterns  
**Domain experiment**: Add `GenerateCode(pattern string) string` that returns AI-style code snippets  
**Next failing test**: GenerateCode("for i in range(5): print(i)") = "for i in range(5): print(i)", want "for i in range(5): print(i)" (pure function, no I/O)

## 3. Average Is All You Need
**Signal**: Data analysis simplicity  
**Domain experiment**: Add `Average(nums []int) int` that returns mean (floor)  
**Next failing test**: Average([]int{1,2,3,4,5}) = 3, want 3 (current `Repeat`/`Replace` handle numbers, but not averages)

## 4. FIM – Linux framebuffer image viewer
**Signal**: Image processing via string manipulation  
**Domain experiment**: Add `DecodeImage(data string) []byte` that simulates image decoding  
**Next failing test**: DecodeImage("PNG") = [80, 78, 71, 13, 10, 26, 10], want [80, 78, 71, 13, 10, 26, 10] (pure function, no I/O)

## 5. A Python Interpreter Written in Python
**Signal**: Self-hosted language interpreters  
**Domain experiment**: Add `Eval(code string) string` that returns evaluated code (mocked)  
**Next failing test**: Eval("1+1") = "2", want "2" (pure function, no I/O)
