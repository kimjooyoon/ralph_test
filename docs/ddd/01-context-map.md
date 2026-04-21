# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, surrogate pairs, and basic string operations (Split, Reverse, Join, Trim, etc.)
- **Unicode Handling**: Focuses on proper decoding/encoding of multi-byte characters, surrogate pairs, and combining marks (Split, Reverse, EncodeBase64)
- **Numeric Operations**: Includes range generation, parsing, and basic arithmetic (Range, Average)
- **Pattern Matching**: Implements regex-like wildcards and substring searches (Match, ContainsWildcard)
- **Encoding/Decoding**: Base64, Hex, and mock encryption/decryption (EncodeBase64, Encode