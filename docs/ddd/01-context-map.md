# Context Map

## Bounded Context: String Manipulation
- **Primary responsibility**: UTF-8 byte-level string operations
- **Key aggregates**:
  - `String` (immutable value object)
- **Domain events**:
  - `SplitEvent` (when string is split by separator)
- **Ubiquitous language**:
  - "Split by empty string" → returns UTF-8 byte array
  - "Multi-byte character" → requires byte-level handling
  - "Code point" → distinct from byte-level operations

## Bounded Context: Data Encoding
- **Primary responsibility**: Base64/Hex encoding/decoding
- **Key aggregates**:
  - `EncodedData` (immutable value object)
- **Domain events**:
  - `EncodeEvent` (when data is encoded)
- **Ubiquitous language**:
  - "Base64 encoding" → strict 6-bit grouping
  - "Surrogate pair" → requires