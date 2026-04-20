# Context Map

## Bounded Contexts
- **String Manipulation**
  - Aggregates: `Split`, `Reverse`, `Join`, `Trim`
  - Invariants: 
    - `Split("中文", "")` returns byte-split Chinese characters `["中", "文"]`
    - `Reverse("\U0001D10D")` treats surrogate pairs as single code points
    - `Join(["中", "文"], "")` reconstructs "中文" from byte-split components

- **Unicode Handling**
  - Aggregates