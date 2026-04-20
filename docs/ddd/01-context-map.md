# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, reverse, encoding, and pattern matching.  
  - **Aggregates**:  
    - `Split` (splits strings into UTF-8 bytes or by separator)  
    - `Reverse` (reverses strings, handling surrogate pairs)  
    - `EncodeBase64