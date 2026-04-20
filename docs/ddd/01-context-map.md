# Context Map — String Manipulation & Unicode

## Bounded Contexts
- **String Manipulation** (primary)
  - Aggregates: `Split`, `Reverse`, `Trim`, `Join`, `Repeat`, `Replace`
  - Domain Events: `StringSplitEvent`, `StringReversedEvent`, `StringTrimmedEvent`
  - Invariants:
    - `Split` must split on UTF-8 bytes for empty separator (ASCII: 1 byte/char, multi-byte