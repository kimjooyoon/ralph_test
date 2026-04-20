# Context Map — String Manipulation & Unicode

## Bounded Contexts
- **String Manipulation** (primary)
  - Aggregate Roots: `Split`, `Reverse`, `Join`, `Trim`, `Repeat`, `Replace`
  - Invariants:
    - `Split(s, "")` returns byte-split strings (UTF-8 bytes, not code points)
    - `Reverse` handles surrogate pairs as single code points
    - `Join` preserves byte boundaries during concatenation
    - `Trim` removes whitespace bytes (ASCII 32, 9-13)
    - `Repeat` ensures byte-level duplication (not