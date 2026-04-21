# Context Map

## Bounded Contexts
- **String Manipulation**: Handles splitting, reversing, encoding, and pattern matching (e.g., `Split`, `Reverse`, `EncodeBase64`)
- **Unicode Processing**: Specializes in multi-byte character handling (e.g., Chinese, emoji, surrogate pairs)
- **Data Transformation**: Focuses on pure functions for data conversion (e.g., `Range`, `Average`, `ParseInt`)
- **Mock I/O**: Simulates I/O operations for testing (e.g., `Log`, `Eval`, `DecodeImage`)

## Aggregate Boundaries
- `Split(s, sep)` treats `sep` as a UTF-8 byte sequence, not a Unicode code point
- `Reverse(s)` must handle surrogate pairs as single code points
- `GenerateCode(pattern)` is a pure function with no side effects