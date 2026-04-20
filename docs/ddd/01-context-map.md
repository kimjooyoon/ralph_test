# Context Map

## Bounded Contexts
- **String Processing** (primary context for `Split`, `Reverse`, `Join`, `Trim`, `Repeat`, `Replace`)
  - Aggregate: `StringHelpers`
  - Invariants:
    - `Split(s string, sep string) []string` must split UTF-8 bytes (not code points) for multi-byte characters
    - `Reverse(s string) string` must handle surrogate pairs and combining marks as single code points
    - `Trim(s string, cutset string) string` must remove all runes in `cutset` from start/end
    - `Repeat(s string, count int) string` must return `s` repeated `count