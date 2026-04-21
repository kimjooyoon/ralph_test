# Context Map

## Bounded Contexts

- **String Manipulation**: Handles UTF-8 byte splitting, surrogate pairs, and Unicode code points.  
  - Aggregates: `Split`, `Reverse`, `Trim`  
  - Invariants:  
    - `Split` must treat surrogate pairs as single code points  
    - `Reverse` must preserve combining marks  
    - UTF-8 validation is required for all inputs  

- **Numeric Operations**: Provides range generation, parsing, and basic math.  
  - Aggregates: `Range`, `ParseInt`, `Average`  
  - Invariants:  
    - `Range` must include inclusive bounds  
    - `ParseInt` must return errors for invalid input  
    - `Average` must floor the mean  

- **Encoding/Decoding**: Implements base64, hex, and mock encryption.  
  - Aggregates: `EncodeBase64`, `EncodeHex`, `EncryptAES`  
  - Invariants:  
    - Encodings must handle UTF-8 input correctly  
    - Mock encryption must return fixed outputs  

## Open Questions
- How to handle ambiguous Unicode code points in `Split`?  
- Should `Average` handle empty slices?  
- Should `EncryptAES` accept key lengths?  
