# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, reversing, and surrogate pair handling. Aggregates: `Split`, `Reverse`, `Join`.
- **Data Analysis**: Provides numeric range generation and average calculation. Aggregates: `Range`, `Average`.
- **Encoding/Decoding**: Implements Base64, Hex, and JSON/YAML serialization. Aggregates: `EncodeBase64`, `EncodeHex`, `SerializeJSON`.
- **AI Code Generation**: Generates code snippets and evaluates expressions. Aggregates: `GenerateCode`, `Eval`.

## Aggregate Boundaries
- `Split` must treat UTF-8 bytes as atomic units for multi-byte characters (e.g., Chinese, emoji).
- `Reverse`