# Context Map

## Bounded Contexts

- **String Manipulation**
  - **Boundaries**: `Split`, `Reverse`, `EncodeBase64`, `ParseInt`, `Range`, `Match`, `ContainsWildcard`, `SerializeJSON`, `EncryptAES`
  - **Aggregates**:
    - `Split`: Splits strings into UTF-8 byte sequences (ASCII: 1 byte/char, UTF-8: multi-byte chars split by byte)
    - `Reverse`: Reverses byte sequences while preserving UTF-8 validity
    - `EncodeBase64`: Converts byte sequences to Base64 strings
    - `ParseInt`: Converts string representations of integers to `