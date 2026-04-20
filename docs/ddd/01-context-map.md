# Context Map

## Bounded Contexts
- **String Manipulation**: Handles splitting, reversing, encoding, and pattern matching (includes `Split`, `Reverse`, `EncodeBase64`, `Match`, `Replace`)
- **Unicode Processing**: Specializes in UTF-8 byte handling for multi-byte characters (e.g., Chinese, emoji) via `Split` and `Reverse`
- **Data Encoding**: Focuses on base64 and hex encoding/decoding (`EncodeBase64`, `EncodeHex`)

## Aggregate Boundaries
- `Split` aggregates: UTF-8 byte splitting for non-empty strings, returns byte slices for multi-byte characters
- `Reverse` aggregates: UTF-8 byte reversal with surrogate pair handling
- `EncodeBase64` aggregates: Base64 encoding with proper padding and character set

## Domain Events/Invariants
- `Split("中文", "")` must return ["中", "文"] (UTF-8 byte splitting)
- `EncodeBase64([]byte("hello"))` must produce "aGVsbG8="
- `Reverse("\U0001D10D")` must treat surrogate pairs as single code points
