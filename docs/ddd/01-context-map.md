# Context Map

## Bounded Contexts
- **DSL Helpers**: Provides string manipulation, encoding, and data processing functions (Split, Reverse, EncodeBase64, etc.)

## Aggregates & Boundaries
- **String Manipulation**: 
  - `Split(s, sep)` splits strings by UTF-8 bytes (when sep is empty)
  - `Reverse(s)` reverses UTF-8 byte sequences, handling surrogate pairs
- **Encoding**: 
  - `EncodeBase64(data)` converts byte slices to base64 strings
  - `EncodeHex(data)` converts bytes to hexadecimal strings
- **Numeric Helpers**: 
  - `