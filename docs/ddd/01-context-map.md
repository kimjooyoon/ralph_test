# Context Map

## Bounded Contexts
- **String Manipulation**
  - Aggregates: `StringSplitter`, `StringReverser`, `StringMatcher`
  - Domain Events: `StringSplitComplete`, `StringReversed`, `PatternMatched`
  - Invariants: 
    - `Split` must split UTF-8 bytes for multi-byte characters
    - `Reverse` must handle surrogate pairs as single code points
    - `Match` must validate regex-like patterns against strings

- **Numeric Operations**
  - Aggregates: `NumericRangeGenerator`, `IntegerParser`
  - Domain Events: `RangeGenerated`, `IntegerParsed`
  - Invariants: 
    - `Range` must generate inclusive numeric sequences
    - `ParseInt` must fail for invalid numeric strings

- **Unicode Handling**
  - Aggregates: `UnicodeValidator`, `SurrogatePairHandler`
  - Domain Events: `UnicodeValidationComplete`, `SurrogatePairProcessed`
  - Invariants: 
    - `Split` must treat surrogate pairs as single code points
    - `Reverse` must preserve combining marks

## Context Boundaries
- `Split` boundary: UTF-8 byte vs Unicode code point splitting
- `Reverse` boundary: single code point vs surrogate pair handling
- `Match` boundary: regex syntax vs pattern matching rules

## Open Questions
- Should `Reverse` handle combining marks as separate code points?
- How to represent surrogate pairs in test cases?
