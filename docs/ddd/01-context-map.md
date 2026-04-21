# Context Map

## Bounded Contexts
1. **String Manipulation**
   - Aggregates: `Split`, `Join`, `Reverse`, `Trim`
   - Invariants: 
     - `Split` treats UTF-8 bytes as atomic units (not code points)
     - `Reverse` preserves surrogate pairs as single code points
     - `Trim` removes whitespace per Unicode definition

2. **Encoding/Decoding**
   - Aggregates: `EncodeBase64`, `EncodeHex`, `DecodeImage`
   - Invariants: 
     - `EncodeHex` must handle UTF-8 input as byte sequences
     - `DecodeImage` returns raw byte arrays for mock processing

3.