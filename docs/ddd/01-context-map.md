# Context Map

## Bounded Contexts
1. **String Manipulation** (core)
   - Handles UTF-8 byte splitting, surrogate pairs, and Unicode normalization
   - Aggregates: `Split`, `Reverse`, `Trim`, `Join`, `Repeat`
   - Invariants: 
     - `Split` with empty separator splits by UTF-8 bytes (not code points)
     - `Reverse` preserves surrogate pairs as single code points
     - `Trim` removes whitespace per Unicode definition

2. **Code Generation** (codex)
   - Transforms natural language patterns into code snippets
   - Aggregates: `GenerateCode`
   - Invariants:
     - Output must be syntactically valid Python code
     - Preserves input patterns as base for generation

3. **Numeric Operations** (range/average)
   - Handles inclusive ranges and mean calculations
   - Aggregates: `Range