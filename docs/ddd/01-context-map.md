# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, surrogate pairs, and Unicode code points
- **Data Encoding**: Base64, hex, and JSON serialization
- **Pattern Matching**: Regex-like wildcards and substring searches
- **Numeric Operations**: Range generation and integer parsing

## Aggregate Boundaries
- `Split` operates on UTF-8 byte sequences, treating surrogate pairs as single code points
- `Reverse` must preserve Unicode normalization for combining marks
- `EncodeBase64` handles byte arrays as opaque data

## Domain Events
- `StringSplitEvent`: Triggered when UTF-8 byte splitting completes
- `UnicodeDecoded`: When surrogate pairs are properly normalized
- `EncodingComplete`: When base64/hex conversion finishes

## Invariants
- UTF-8 decoding must fail on invalid byte sequences
- Empty separator splits by byte (ASCII) or code point (Unicode)
- Surrogate pairs must be treated as single code points
