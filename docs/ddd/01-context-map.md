# Context Map

## Bounded Contexts
- **String Manipulation Helpers**: Handles string operations with UTF-8 byte and code point awareness.  
  - **Aggregate Boundaries**:  
    - `Split` (splits by UTF-8 bytes, not code points)  
    - `Reverse` (handles surrogate pairs and combining marks)  
    - `EncodeBase64` (encodes byte slices to base64)  
    - `Match` (regex-like pattern matching)  
    - `Average` (numeric range and averaging)  

## Domain Events/