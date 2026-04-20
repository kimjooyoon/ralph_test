# Context Map

## Bounded Contexts
- **String Manipulation**
  - Aggregates: `Split`, `Join`, `Trim`, `Reverse`
  - Invariants: 
    - `Split("", "")` returns empty slice
    - `Split(s, "")` splits by UTF-8 byte (not Unicode code point)
    - `Split("中文", "")` returns ["中", "文"]
    - `Split("\U0001D10D", "")` returns ["\U0001D10D"]
  - Domain Events: `StringSplitEvent`, `String