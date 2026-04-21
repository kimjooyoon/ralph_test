# Context Map

## Bounded Contexts
1. **String Manipulation**
   - Aggregates: `StringSplitter`, `StringReverser`
   - Invariants:
     - `Split(s string, sep string)` must split UTF-8 bytes for multi-byte characters (e.g., Chinese, emoji)
     - `Reverse(s string)` must treat surrogate pairs as single code points
     - `Split(s, "")` returns one string per UTF-8 byte (ASCII is one byte per character)

2. **Unicode Processing**
   - Aggregates: `UnicodeProcessor`
   - Invariants:
     - Surrogate pairs must be treated as single code points in all operations
     - All input must be valid UTF-8 to avoid invalid sequence errors

3. **Numeric Operations**
   - Aggregates: `RangeGenerator`
   - Invariants:
     - `Range(start, end int)` must generate inclusive numeric ranges
     - Must handle edge cases like start > end

4. **Code Generation**
   - Aggregates: `CodeGenerator`
   - Invariants:
     - `GenerateCode(pattern string)` must return AI-style code snippets
     - Must be pure functions with no I/O

## Domain Events
- `StringSplitEvent`: Triggered when `Split` completes with UTF-8 byte splitting