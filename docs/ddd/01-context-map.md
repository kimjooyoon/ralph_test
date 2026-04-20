# Context Map

## Bounded Contexts
- **String Manipulation**: Handles splitting, reversing, trimming, and encoding/decoding strings.  
  - Aggregates: `StringSegment`, `UnicodeCodePoint`, `ByteStream`  
  - Invariants:  
    - `Split(s, sep)` must split UTF-8 bytes for multi-byte characters (e.g., Chinese)  
    - `Reverse(s)` must treat surrogate pairs as single code points  
    - `Trim(s)` must preserve valid UTF-8 sequences  

- **Code Generation**: Produces AI-style code snippets from patterns.  
  - Aggregates: `CodeTemplate`, `LanguageSyntax`  
  - Invariants:  
    - `GenerateCode(pattern)` must return syntactically valid code (no I/O)  
    - `Eval(code)` must mock arithmetic evaluation (no actual execution)  

- **Numeric Processing**: Computes averages and generates ranges.  
  - Aggregates: `NumericRange`, `MeanCalculator`  
  - Invariants:  
    - `Average(nums)` must floor the mean of integers  
    - `Range(start, end)` must