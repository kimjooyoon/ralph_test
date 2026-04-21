# Context Map

## Bounded Contexts

### 1. String Manipulation
- **Responsibility**: Core string operations (split, reverse, trim, etc.)
- **Aggregates**:
  - `Split` (UTF-8 byte splitting, surrogate pair handling)
  - `Reverse` (surrogate pair & combining mark awareness)
  - `Trim` (whitespace removal with Unicode-aware logic)
- **Invariants**:
  - `Split("", s)` returns byte-per-UTF-8 (not