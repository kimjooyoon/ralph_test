# Context Map — String Manipulation

## Bounded Context: String Manipulation
- **Primary responsibility**: UTF-8 byte-level string operations with Unicode awareness
- **Key aggregates**:
  - `StringSplitter`: Handles byte-wise splitting with UTF-8 decoding
  - `UnicodeDecoder`: Manages surrogate pairs and combining marks
- **Domain events**:
  - `ByteSplitComplete`: When a string is split into UTF-8 bytes
  - `UnicodeDecoded`: When a multi-byte sequence is properly decoded
- **Ubiquitous language**:
  - "Split by byte" (vs. "split by code point")
  -