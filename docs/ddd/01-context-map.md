# Context Map — dsl-maker

## Bounded Contexts

- **String Manipulation**: Handles UTF-8, Unicode, and byte-level operations (Split, Reverse, Trim, Join, Replace, Repeat)
- **Code Generation**: AI-style code snippet creation (GenerateCode)
- **Numeric Helpers**: Range generation, parsing, and basic math (Range, Average)
- **Encoding/Decoding**: Base64, Hex, and other data transformations (EncodeBase64, EncodeHex)
- **Pattern Matching**: Regex-like wildcards and substring checks (Match, ContainsWildcard)
- **Serialization**: JSON and YAML formatting (SerializeJSON, SerializeYAML)
- **Logging/Debugging**: Mock logging and timestamp generation (Log, Timestamp)

## Aggregate Boundaries

- **String Manipulation Aggregate**: Enforces UTF-8 byte/character splitting, surrogate pair handling, and Unicode normalization
- **Code Generation Aggregate**: Ensures AI-style code snippets are pure functions with no I/O
- **Numeric Aggregate**: Maintains numeric range validity and integer parsing constraints
- **Encoding Aggregate**: Guarantees valid base64/hex encoding without external dependencies

## Domain Events

- `StringSplitEvent`: Triggered when UTF-8 byte splitting completes
- `CodeGeneratedEvent`: Fired after AI