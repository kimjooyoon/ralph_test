# Context Map

## Bounded Contexts
- **String Manipulation**: Handles core string operations (Split, Reverse, Join, Trim, etc.)
  - Aggregates: StringSegment, CodeSnippet
  - Invariants: 
    - `Split(s, sep)` must split UTF-8 bytes for multi-byte characters
    - `Reverse(s)` must handle surrogate pairs as single code points
- **Code Generation**: AI-style code snippet creation
  - Aggregates: CodeTemplate, GeneratedCode
  - Invariants: 
    - `GenerateCode(pattern)` must return valid syntax without I/O
    - `Eval(code)` must mock execution of simple expressions
- **Unicode Processing**: Specialized handling of Unicode edge cases
  - Aggregates: UnicodeSegment, SurrogatePair
  - Invariants: 
    - `Split(s, "")` must split by UTF-8 bytes (not code points)
    - `TestSplitSurrogatePairs()` must validate emoji handling

## Context Boundaries
- `Split` operates on byte-level UTF-8 sequences (not code points)
- `Reverse` treats surrogate pairs as single code points
- `GenerateCode` is a pure function