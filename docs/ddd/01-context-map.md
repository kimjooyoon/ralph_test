# Context Map

## Bounded Contexts
- **String Manipulation**: Handles basic string operations (split, reverse, join, trim)
- **Unicode Handling**: Specializes in multi-byte character processing (Chinese, emoji, surrogate pairs)
- **Encoding/Decoding**: Manages binary-to-text encoding (Base64, Hex)
- **Codex AI**: Implements AI-style code generation patterns
- **Numeric Operations**: Provides range generation and numeric checks

## Aggregate Boundaries
- **String Manipulation**:
  - `Split(s, sep)` - Splits strings by UTF-8 bytes (for multi-byte characters)
  - `Reverse(s)` - Reverses string while preserving surrogate pairs
  - `Join(sep, parts)` - Concaten