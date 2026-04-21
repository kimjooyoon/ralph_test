# Context Map & Domain Model

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 string operations (Split, Reverse, EncodeHex)
- **Numeric Operations**: Range generation, integer parsing
- **Pattern Matching**: Regex-like wildcards
- **Encoding/Decoding**: Base64, Hex, Logging

## Aggregate Boundaries
- **String Manipulation Aggregate**:
  - `Split(s, sep)` handles UTF-8 byte sequences (surrogate pairs as single code points)
  - `Reverse(s)` preserves surrogate pairs and combining marks
  - `EncodeHex(data)` processes UTF-8