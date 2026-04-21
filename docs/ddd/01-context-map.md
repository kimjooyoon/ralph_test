# Context Map

## Bounded Contexts
- **String Manipulation**
  - Responsible for: UTF-8 byte splitting, surrogate pair handling, and pure string operations
  - Aggregates:
    - `Split` (splits strings by UTF-8 bytes)
    - `Reverse` (handles surrogate pairs and combining marks)
    - `Join` (combines strings with separators)
  - Invariants:
    - `Split("")` returns empty slice
    - `Split(s, "")` returns UTF-8 byte array
    - Surrogate pairs are treated as single code points
    - All operations are pure functions

## Context Boundaries
- `Split` boundary: 
  - Input: UTF-8 string and separator
  - Output: Slice of UTF-8 bytes (for empty separator)
  - Violates boundary: I/O operations
- `Reverse` boundary:
  - Input: UTF-8 string
  - Output: Reversed string with proper surrogate pair handling
  - Violates boundary: Network operations

## Domain Events
- `StringSplitEvent`: