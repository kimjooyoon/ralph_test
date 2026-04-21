# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, Unicode reversal, and surrogate pairs (e.g., `Split`, `Reverse`, `Trim`)
- **Data Generation**: Produces AI-style code snippets, numeric ranges, and mock evaluations (e.g., `GenerateCode`, `Range`, `Eval`)
- **Encoding/Decoding**: Base64, hexadecimal, and JSON/YAML serialization (e.g., `EncodeBase64`, `EncodeHex`, `SerializeJSON`)
- **Logging/Time**: Mock logging and timestamp generation (e.g., `Log`, `Timestamp`)

## Aggregate Boundaries
- **String Manipulation Aggregate**: Enforces UTF-8 byte splitting rules for multi-byte characters (e.g., Chinese, emoji)
- **Data Generation Aggregate**: Ensures AI-style code generation adheres to pure function constraints (no I/O, no side effects)
- **Encoding Aggregate**: Validates base64/hex encoding rules for valid UTF-8 input
- **Logging Aggregate**: Treats logging as a pure function for mock debugging scenarios

## Domain Events/Invariants
- **UTF-8 Splitting Invariant**: `Split(s, "")` must return one string per UTF-8 byte for