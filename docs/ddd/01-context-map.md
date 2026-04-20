# Context Map

## Bounded Contexts
- **String Manipulation**: Handles splitting, joining, reversing, trimming, and basic pattern matching (e.g., `Split`, `Reverse`, `Trim`)
- **Encoding/Decoding**: Base64, hexadecimal, and mock image/JSON/YAML serialization (e.g., `EncodeBase64`, `EncodeHex`, `SerializeJSON`)
- **Unicode Handling**: Specialized string operations for multi-byte characters, surrogate pairs, and combining marks (e.g., `Split` with UTF-8 byte splitting, `Reverse` for emoji