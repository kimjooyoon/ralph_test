# Context Map

## Bounded Contexts
- **String Manipulation**: Handles splitting, reversing, trimming, joining, and pattern matching (Split, Reverse, Trim, Join, Match, ContainsWildcard)
- **Encoding/Decoding**: Base64, Hex, JSON, YAML, and mock encryption (EncodeBase64, EncodeHex, SerializeJSON, SerializeYAML, EncryptAES)
- **Numeric Operations**: Range generation, averaging, and parsing (Range, Average, ParseInt)
- **AI Code Generation**: Generates mock code snippets and evaluates expressions (GenerateCode, Eval)
- **Unicode Handling**: Special focus on UTF-8 byte splitting, surrogate pairs, and combining marks (Split, Reverse)

## Aggregate Boundaries
- **String Manipulation Aggregate**: Enforces pure, stateless operations on strings (no I/O, no side effects)
- **Encoding Aggregate**: Ensures data transformations are reversible and lossless
- **Numeric Aggregate**: Maintains mathematical invariants for arithmetic operations
- **AI Code Aggregate**: Produces deterministic outputs based on input patterns

## Domain Events
- `StringSplitEvent`: Triggered when a string is split by UTF-8 bytes
- `UnicodeReverseEvent`: Occurs when a string with surrogate pairs is reversed
- `Base64EncodeEvent`: Fired during successful base64 encoding
- `CodeGeneratedEvent`: Not