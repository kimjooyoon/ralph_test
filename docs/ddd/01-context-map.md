# Context Map

## Bounded Contexts
- **String Manipulation Helpers**: Contains all string-related operations (Split, Reverse, EncodeBase64, etc.)

## Aggregates & Invariants
- **Split**: 
  - Invariant: `Split(s, "")` returns a slice of UTF-8 bytes (each byte as a string)
  - Test: `Split("中文", "")` must return `["中", "文"]` (byte-level split)
- **Reverse**: 
  - Invariant: Handles surrogate pairs and combining marks as single code points
  - Test: `Reverse("\U0