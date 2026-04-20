# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, reversing, and pattern matching (Split, Reverse, Match, ContainsWildcard)
- **Encoding/Decoding**: Base64, hexadecimal, and JSON/YAML serialization (EncodeBase64, EncodeHex, SerializeJSON, SerializeYAML)
- **Numeric Operations**: Range generation, averaging, and integer parsing (Range, Average, ParseInt)
- **AI Code Generation**: Generates code snippets and evaluates expressions (GenerateCode, Eval)
- **Image Processing**: Simulates image decoding from string data (DecodeImage)

## Aggregate Boundaries
- **String Manipulation Aggregate**: 
  - `Split(s string, sep string) []string` - Splits strings by UTF-8 bytes (when sep is empty)
  - `Reverse(s string) string` - Reverses UTF-8 byte sequences
  - `Match(pattern, s string) bool` - Basic regex-like pattern matching
  - `ContainsWildcard(s, pattern string) bool` - Wildcard substring search

- **Encoding Aggregate**:
  - `EncodeBase64(data []byte) string` - Base64 encoding
  - `EncodeHex(data []byte) string`