# Context Map

## Bounded Contexts
- **String Manipulation**
  - Aggregates: `Split`, `Join`, `Trim`, `Reverse`, `Replace`, `Repeat`
  - Invariants: 
    - `Split` must split on UTF-8 bytes (not code points)
    - `Reverse` must reverse bytes, not code points
    - `Replace` must handle multi-byte sequences
- **Unicode Handling**
  - Aggregates: `Split`, `Reverse`, `EncodeBase64`
  - Invariants: 
    - Surrogate pairs treated as single code points
    - Valid UTF-8 input required for all operations
- **Numeric Operations**
  - Aggregates: `Range`, `ParseInt`, `Average`
  - Invariants: 
    - `Range` generates inclusive sequences
    - `ParseInt` returns error for invalid input
    - `Average` floors to nearest integer
- **Regex-Like Matching**
  - Aggregates: `Match`, `ContainsWildcard`
  - Invariants: 
    - `Match` supports basic regex syntax