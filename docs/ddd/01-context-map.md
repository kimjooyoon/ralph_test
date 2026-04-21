# Context Map

## Bounded Contexts
- **String Processing** (primary context for `Split`, `Join`, `Reverse`, `Trim`, `ContainsWildcard`)
- **Encoding/Decoding** (for `EncodeBase64`, `EncodeHex`, `DecodeImage`)
- **Numeric Operations** (for `Range`, `ParseInt`, `Average`)
- **Regex-Like Matching** (for `Match`, `ContainsWildcard`)
- **Logging/Debugging** (for `Log`, `Timestamp`)

## Aggregate Boundaries
- `Split` operates on **UTF-8 byte sequences** (not code points) in the String Processing context
- `EncodeHex` handles **UTF-8 input as byte sequences** in the Encoding context
- `Range` generates **inclusive integer sequences** in Numeric Operations
- `Match` and `ContainsWildcard` work with **pattern matching rules** in Regex-Like Matching