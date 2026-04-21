# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, reverse, and surrogate pair processing. Core functions: `Split`, `Reverse`, `EncodeBase64`.
- **Unicode Handling**: Focuses on multi-byte character support (e.g., Chinese, emoji). Invariants: surrogate pair integrity, combining marks, UTF-8 validation.
- **Encoding/Decoding**: Pure functions for data transformation (base64, hex). No I/O, no side effects