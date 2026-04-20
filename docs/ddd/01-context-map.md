# Context Map

## Bounded Contexts
- **String Manipulation**: Contains `Split`, `Join`, `Trim`, `Reverse`, `Replace`, `Repeat`, `Echo` (core string operations)
- **Unicode Handling**: Specialized for surrogate pairs, combining marks, and multi-byte characters (e.g., Chinese, emoji)
- **Data Analysis**: Includes `Average`, `Range`, `ParseInt` (numeric helpers)
- **Encoding/Decoding**: Base64, Hex, JSON, YAML, AES (serialization/security)
- **Logging/Debugging**: Timestamp generation, mock logging

## Aggregate Boundaries
- `Split` aggregate: Handles UTF-8 byte splitting, surrogate pairs, and multi-byte characters
- `Reverse` aggregate: Manages combining marks and surrogate pairs
- `Replace`/`Repeat` aggregate: Pure functions with no side effects
- `ParseInt` aggregate: Validates numeric input and handles errors

## Domain Events/Invariants
- `Split` must split UTF-8 bytes for multi-byte characters (e.g., "中文" → ["中", "文"])
- `Reverse` must preserve combining marks and surrogate pairs
- `ParseInt` must return error for invalid numeric strings
- `Average` must floor the mean of a slice of integers

## Open Questions
- Should `EncodeBase64`/`EncodeHex` belong in a separate bounded context?
- How to handle invalid UTF-8 input in `Split`? (Currently, tests use valid UTF-8)
- Should `Eval`/`DecodeImage` be in a separate context for security reasons?
