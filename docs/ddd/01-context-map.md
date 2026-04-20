# Context Map

## Bounded Contexts
- **String Manipulation**: Handles core string operations (split, join, trim, reverse)
  - Aggregates: `Split`, `Join`, `Trim`, `Reverse`
  - Invariants: 
    - `Split` must split UTF-8 bytes for multi-byte characters
    - `Reverse` handles surrogate pairs and combining marks
- **Pattern Matching**: Implements regex-like pattern matching
  - Aggregates: `Match`, `ContainsWildcard`
  - Invariants: 
    - `Match` validates regex-like syntax
    - `ContainsWildcard` supports * and ? wildcards
- **Numeric Processing**: Manages number ranges and conversions
  - Aggregates: `Range`, `ParseInt`
  - Invariants: 
    - `Range` generates inclusive numeric sequences
    - `ParseInt` validates string-to-integer conversions
- **Encoding/Decoding**: Provides base64 and hex encoding
  - Aggregates: `EncodeBase64`, `EncodeHex