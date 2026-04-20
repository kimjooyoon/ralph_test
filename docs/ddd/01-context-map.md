# Context Map

## Bounded Contexts
1. **String Manipulation** (primary context)
   - Aggregates: `StringOperations`
   - Invariants: 
     - `Split` must split UTF-8 bytes for multi-byte characters
     - `Reverse` must handle surrogate pairs and combining marks
     - `Match` must reject non-numeric strings with regex-like pattern

2. **Unicode Handling** (sub-context of String Manipulation)
   - Aggregates: `UnicodeProcessor`
   - Invariants:
     - `Split` must treat surrogate pairs as single code