# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 splitting, reversal, and pattern matching (Split, Reverse, Match)
- **Encoding/Decoding**: Base64, hex, and image decoding (EncodeBase64, DecodeImage)
- **Numeric Operations**: Range generation, parsing, and averaging (Range, ParseInt, Average)
- **Regex-Like Patterns**: Wildcard matching and simple pattern validation (Match, ContainsWildcard)

## Aggregate Boundaries
- **String Manipulation Aggregate**: 
  - Split: UTF-8 byte splitting for multi-byte characters
  - Reverse: Surrogate pair preservation
  - Match: Regex-like pattern validation
- **Encoding Aggregate**:
  - EncodeBase64: Pure byte-to-string encoding
  - DecodeImage: Simulated image decoding from byte slices

## Domain Events
- `StringSplitEvent`: Triggered when UTF-8 splitting completes
- `EncodingComplete`: Fired after successful base64 encoding
- `RangeGenerated`: Notified when numeric ranges are created