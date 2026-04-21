# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 decoding, splitting, reversing, and encoding (Split, Reverse, EncodeBase64)
- **Numeric Operations**: Range generation, parsing, and averaging (Range, ParseInt, Average)
- **Pattern Matching**: Regex-like wildcards and substring checks (Match, ContainsWildcard)
- **Data Encoding**: Base64, hexadecimal, and serialization formats (EncodeBase64, EncodeHex, SerializeJSON)

## Aggregates
- **String Manipulation Aggregate**: Ensures UTF-8 safety for Split (byte-level)