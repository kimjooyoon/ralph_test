# Context Map

## Bounded Contexts
- **String Manipulation**: Handles splitting, reversing, trimming, joining, and pattern matching (Split, Reverse, Trim, Join, Match, ContainsWildcard)
- **Encoding/Decoding**: Base64, Hex, and image decoding (EncodeBase64, EncodeHex, DecodeImage)
- **Numeric Operations**: Range generation, averaging, and parsing (Range, Average, ParseInt)
- **AI Code Generation**: Generates code snippets from patterns (GenerateCode)
- **Unicode Handling**: Specialized splitting for multi-byte characters and surrogate pairs (Split with UTF-8 byte splitting)

## Aggregate Boundaries
- **String