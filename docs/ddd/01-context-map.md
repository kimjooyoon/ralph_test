# Context Map

## Bounded Contexts
- **String Manipulation DSL**: Pure string operations with UTF-8/Unicode awareness

## Aggregates & Invariants
- **Split**:
  - `sep == ""` → split by UTF-8 byte (not Unicode code point)
  - Must handle