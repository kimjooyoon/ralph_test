# Context Map

## Bounded Contexts

### 1. String Manipulation
- **Aggregate Boundaries**: 
  - `Split(s, sep)` splits strings by UTF-8 bytes (for `sep == ""`)
  - `Reverse(s)` reverses string bytes
  - `Trim(s)` removes whitespace from start/end
- **Domain Events**:
  - `StringSplitEvent` (when split by byte)
  - `StringReverseEvent` (when reversed)
  - `StringTrimEvent` (when whitespace removed)
- **Invariants**:
  - `Split("中文", "")` → `["中", "文"]` (byte-split Chinese)