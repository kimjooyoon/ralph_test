# Context Map

## Bounded Contexts
1. **StringHelpers**
   - Aggregates: `Split`, `Reverse`, `Join`, `Trim`, `Replace`, `Repeat`, `ContainsWildcard`, `Match`
   - Invariants:
     - `Split` must treat UTF-8 byte sequences as atomic units for surrogate pairs (e.g., emoji)
     - `Reverse` must preserve surrogate pairs as single code points
     - `Match` must support regex-like syntax for pattern matching
     - `ContainsWildcard` must handle `*` and `?` wildcards