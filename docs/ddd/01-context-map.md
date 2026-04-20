# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, reverse, and encoding/decoding.  
  - `Split` (UTF-8 byte splitting for multi-byte chars)  
  - `Reverse` (surrogate pairs, combining marks)  
  - `EncodeBase64` (pure