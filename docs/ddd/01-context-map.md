# Context Map

## Bounded Contexts
1. **String Manipulation**
   - Aggregate Roots: `Split`, `Reverse`, `Join`
   - Invariants:
     - `Split(s, sep)` returns UTF-8 byte splits for `sep == ""` (Chinese characters, surrogate pairs)
     - `Reverse(s)` preserves surrogate pairs and combining marks
     - `Join(sep, parts)` maintains UTF-8 byte integrity

2. **Unicode Handling**
   - Aggregate Roots: `Split`, `Reverse`, `EncodeBase64`
   - Invariants:
     - Surrogate pairs treated as single code points
     - Valid UTF-8 input required for all operations
     - Proper byte stream handling for multi-byte characters

3. **Data Conversion**
   - Aggregate Roots: `ParseInt`, `Range`, `Average`
   - Invariants:
     - `ParseInt(s)` validates