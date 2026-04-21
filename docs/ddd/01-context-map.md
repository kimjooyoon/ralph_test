# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 string operations (Split, Reverse, Join, Trim)
  - Aggregate Roots: `StringSegment`, `UnicodeCodePoint`
  - Invariants: 
    - `Split(s string, sep string) []string` must split by UTF-8 bytes when `sep == ""`
    - `Reverse(s string) string` must handle surrogate pairs as single code points
    - `Join(sep string, parts []string) string` must preserve UTF-8 integrity
- **Data Encoding**: Base64, Hex, JSON, YAML encoding/decoding
  - Aggregate Roots: `EncodedData`, `BinaryPayload`
  - Invariants: 
    - `EncodeBase64(data []byte) string` must produce valid Base64
    - `DecodeImage(data string) []byte` must simulate image decoding
- **Numeric Operations**: Range generation, parsing, averaging
  - Aggregate Roots: `NumericRange`, `IntegerValue`
  - Invariants: 
    - `Range(start, end int) []int` must generate inclusive ranges
    - `