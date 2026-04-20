# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, surrogate pairs, and combining marks for string operations.

## Aggregates
- **Split**: Splits strings into UTF-8 bytes (for empty separator), ensuring proper handling of multi-byte characters, surrogate pairs, and combining marks.
- **Reverse**: Reverses strings while preserving UTF-8 byte order and handling surrogate pairs correctly.
- **Trim**: Removes leading/trailing UTF-8 bytes while respecting valid UTF-8 sequences.

## Invariants
- `Split(s, "")` returns one string per UTF-8 byte of `s` (ASCII