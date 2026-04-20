# Context Map

## Bounded Contexts
- **String Manipulation** (primary context)
  - Aggregates: `Split`, `Join`, `Trim`, `Reverse`, `Replace`, `Repeat`, `Match`
  - Invariants: 
    - `Split` treats empty separator as UTF-8 byte split (ASCII 1 byte, multi-byte chars split by byte)
    - `Reverse` handles surrogate pairs as single code points
    - `Match` validates regex-like patterns
- **Numeric Operations**
  - Aggregates: `Range`, `ParseInt`
  - Invariants: 
    - `Range(1,5)` returns [1,2,3,4,5]
    - `ParseInt("123")` returns 123
- **Unicode Handling**
  - Aggregates: `Split`, `Reverse`
  - Invariants: 
    - Surrogate pairs treated as single code points
    - Chinese characters split by UTF-8 byte
- **Data Encoding**
  - Aggregates: `EncodeBase64`, `EncodeHex`
  - Invariants: 
    - `EncodeBase64("hello")` returns "aGVsbG8="
    - `EncodeHex("abc")` returns "616263"

## Domain Events
- `StringSplitEvent` (when `Split` completes)
- `StringReversedEvent` (when `Reverse` completes)