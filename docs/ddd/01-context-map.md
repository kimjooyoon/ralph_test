# Context Map

## Bounded Contexts
- **String Manipulation**
  - Aggregate: `Split` (splits UTF-8 bytes, not code points)
  - Invariants: 
    - `Split("中文", "")` → `["中", "文"]` (byte-level split)
    - `Split("\U0001D10D", "")` → `["\U0001D10D"]` (surrogate pair as single unit)
  - Domain Events: 
    - `ByteSplitEvent` (triggered when splitting by UTF-8 bytes)
    - `SurrogatePairDetected` (when handling emoji/extended characters)

- **Unicode Handling**
  - Aggregate: `Reverse` (reverses code points, not byte order)
  - Invariants: 
    - `Reverse("\u0300\u0301")` → `["\