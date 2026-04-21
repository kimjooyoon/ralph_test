# Context Map

## Bounded Contexts
- **String Manipulation**: Split, Reverse, Trim, Join, Replace, Repeat, Echo
- **Encoding/Decoding**: EncodeBase64, EncodeHex, DecodeImage (mocked)
- **Numeric Operations**: Range, Average, ParseInt
- **Pattern Matching**: Match, ContainsWildcard
- **Code Generation**: GenerateCode (new context)
- **Logging/Serialization**: Log, Timestamp, SerializeJSON, SerializeYAML
- **Encryption**: EncryptAES

## Aggregate Boundaries
- **String Manipulation**: All string operations are atomic within this context
- **Encoding/Decoding**: Separate from string manipulation (pure functions)
- **Code Generation**: Isolated from other contexts to avoid side effects

## Domain Events
- `CodeGenerated`: Triggered when `GenerateCode` produces a snippet
- `