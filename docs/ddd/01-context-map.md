# Context Map

## Bounded Contexts
- **String Manipulation**
  - Aggregate: `StringSplitter`
  - Invariants: 
    - `Split("中文", "")` returns UTF-8 byte splits for Chinese characters
    - Surrogate pairs are treated as single code points
    - Empty input returns empty slice

- **AI Code Generation**
  - Aggregate: `CodeGenerator`
  - Invariants: 
    - `GenerateCode("for i in range(5): print(i)")` returns AI-style code snippets
    - Special characters are escaped in output
    - Pure function with no I/O

- **Numeric Operations**
  - Aggregate: `NumericProcessor`
  - In