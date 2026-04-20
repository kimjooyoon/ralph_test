# Context Map

## Bounded Contexts

1. **String Manipulation**
   - Aggregate Roots: `Split`, `Reverse`, `Join`, `Trim`, `Repeat`, `Replace`
   - Invariants:
     - `Split` must split UTF-8 bytes (not code points) for multi-byte characters
     - `Reverse` must handle surrogate pairs and combining marks as single code points
     - `Join` must preserve UTF-8 integrity across concatenation
     - `Trim` must remove whitespace while preserving UTF-8 validity

2. **Unicode Handling**
   - Aggregate Roots: `Reverse`, `Split`, `EncodeBase64`, `DecodeHex`
   - Invariants:
     - Surrogate pairs must be treated as single code points
     - UTF-8 validation