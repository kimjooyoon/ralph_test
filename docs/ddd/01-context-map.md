# Context Map

## Bounded Contexts
- **String Manipulation**
  - Aggregates: `Split`, `Join`, `Trim`, `Reverse`, `Replace`, `Repeat`, `Echo`
  - Domain Events: `StringSplit`, `StringJoin`, `StringTrim`, `StringReverse`, `StringReplace`, `StringRepeat`
  - Invariants: 
    - `Split` must split on UTF-8 bytes (not Unicode code points)
    - `Reverse` must handle surrogate pairs and combining marks
    - `ContainsWildcard` must support `*` and `?` wildcards

- **Encoding/Decoding**
  -