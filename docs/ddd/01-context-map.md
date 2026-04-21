# Context Map

## Bounded Contexts
- **String Manipulation** (primary context)
  - Aggregates: `Split`, `Reverse`, `Join`, `Trim`, `Replace`, `Repeat`
  - Invariants: 
    - `Split` must split UTF-8 byte sequences (not code points)
    - `Reverse` must handle surrogate pairs as single code points
    - `Join` must preserve UTF-8 integrity
  - Domain Events: 
    - `StringSplitEvent` (triggered by `Split`)
    - `StringReversedEvent` (triggered by `Reverse`)
- **Encoding** (secondary context)
  - Aggregates: `EncodeBase64`, `EncodeHex`
  - Invariants: 
    - `EncodeHex` must handle UTF-8 input as