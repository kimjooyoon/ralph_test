# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, surrogate pairs, and Unicode code points (Split, Reverse, Trim)
- **Numeric Operations**: Integer parsing, range generation, and averaging (ParseInt, Range, Average)
- **Pattern Matching**: Regex-like wildcards and substring searches (Match, ContainsWildcard)
- **Encoding/Decoding**: Base64, hexadecimal, and mock encryption (EncodeBase64, EncodeHex, EncryptAES)
- **Serialization**: JSON and YAML formatting (SerializeJSON, SerializeYAML)

## Aggregate Boundaries
- **String Splitter**: Splits strings by UTF-8 bytes or code points (Split)
- **Unicode Handler**: Manages surrogate pairs and combining marks (Reverse, Split)
- **Numeric Parser**: Converts strings to integers and generates ranges (ParseInt, Range)
- **Pattern Matcher**: Validates regex-like patterns and wildcards (Match, ContainsWildcard)
- **Encoder**: Converts data to encoded strings (EncodeBase64, EncodeHex)

## Domain Events/Invariants
- **UTF-8 Splitting**: Must split by byte boundaries for multi-byte characters (Split("中文", ""))
- **Surrogate Pair Handling**: Must treat surrogate pairs as single code points (Split("\U0001D10D", ""))
- **Integer Parsing**: Must return errors for invalid input (ParseInt("abc"))
- **Regex Matching**: Must validate pattern syntax and content (Match("^[0-9]+$", "123abc"))
- **Mock Encryption**: Must return fixed outputs for testing (EncryptAES("hello", "key") = "encrypted")

## Open Questions
- Should `Split` handle empty strings differently across contexts?
- How to differentiate between "code point" and "byte" splitting in documentation?
- Should surrogate pair handling be a separate bounded context?