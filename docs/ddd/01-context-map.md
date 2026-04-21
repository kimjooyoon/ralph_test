# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, surrogate pairs, and Unicode edge cases (Split, Reverse, Trim, Join, etc.)
- **Data Encoding**: Base64, hexadecimal, and binary encoding/decoding (EncodeBase64, EncodeHex)
- **Numeric Operations**: Range generation, integer parsing, and averaging (Range, ParseInt, Average)
- **Pattern Matching**: Regex-like wildcards and substring searches (Match, ContainsWildcard)
- **Mock I/O**: Logging, timestamping, and mock evaluation (Log, Timestamp, Eval)
- **AI Code Generation**: AI-style code snippet creation (GenerateCode)

## Aggregate Boundaries
- **String Manipulation Aggregate**: Ensures UTF-8 byte splitting respects multi-byte characters (e.g., Chinese, emoji)
- **Encoding Aggregate**: Guarantees base64/hex encoding/decoding adheres to RFC standards
- **Numeric Aggregate**: Maintains numerical invariants (e.g