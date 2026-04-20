# Context Map

## Bounded Context: String Manipulation Helpers
- **Primary domain**: String processing with UTF-8 awareness
- **Key aggregates**:
  - `Split(s string, sep string) []string`: UTF-8 byte splitting for multi-byte characters
  - `GenerateCode(pattern string) string`: AI-style code generation
  - `Reverse(s string) string`: Unicode-aware string reversal
  - `Range(start, end int) []int`: Numeric range generation
- **Shared language**:
  - "UTF-8 byte splitting" → explicit byte-level segmentation
  - "AI-style code" → generated code patterns
  - "Surrogate pairs" → Unicode handling for emoji/extended characters
- **Invariants**:
  - `Split` with empty separator returns byte-level segments
  - `GenerateCode` produces valid syntax without I/O
  - `Reverse` correctly handles surrogate pairs

## Open questions
- How to handle invalid UTF-8 sequences in input?
- Should `Split` support different segmentation strategies (e.g., code points vs bytes)?
