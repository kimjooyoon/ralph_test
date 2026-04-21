# Context Map

## Bounded Contexts
- **String Manipulation Helpers**: Handles UTF-8 byte splitting, Unicode reversal, and surrogate pair processing.

## Aggregates & Invariants
- **Split**: 
  - Splits strings by UTF-8 bytes when separator is empty.
  - Must handle multi-byte characters (e.g