# Context Map

## Bounded Contexts
- **String Manipulation**: Handles string splitting, reversing, encoding, and pattern matching. Core operations include `Split`, `Reverse`, `EncodeHex`, `EncodeBase64`, `Match`, and `ContainsWildcard`.
- **Encoding**: Focuses on data encoding/decoding (e.g., Base64, Hex). Includes `EncodeBase64`, `EncodeHex`, and `DecodeImage` (as a mock).
- **Numeric Operations**: Provides range generation and parsing (`Range`, `ParseInt`).
- **Logging/Serialization**: Mock logging and data format conversion (`Log`, `Timestamp`, `SerializeJSON`, `SerializeYAML`).

## Aggregate Boundaries
- **String Manipulation Aggregate**: 
  - `Split(s, sep)` splits UTF-8 bytes (not code points) for multi-byte characters.
  - `Reverse(s)` handles surrogate pairs and combining marks as single code points.
  - `Match(pattern, s)` and `ContainsWildcard(s, pattern)` enforce regex-like syntax rules.
- **Encoding Aggregate**:
  - `