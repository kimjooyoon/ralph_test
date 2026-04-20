# Context Map

## Bounded Contexts
- **String Manipulation**
  - Primary responsibility: UTF-8 byte-level string operations
  - Key functions: Split, Reverse, EncodeBase64, Range, ParseInt
  - Invariants:
    - `Split("", "")` returns empty slice
    - `Split("中文", "")` returns ["中", "文"] (UTF-8 byte split)
    - `Reverse("\U0001D10D")` returns same emoji (surrogate pair handling)
    - `EncodeBase64