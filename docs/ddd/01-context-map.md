# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, surrogate pairs, and Unicode normalization.  
  - Functions: `Split`, `Reverse`, `EncodeBase64`, `DecodeImage` (mocked)  
  - Invariants:  
    - `Split("中文", "")` returns ["中", "文"] (byte-level split)  
    - `Reverse("\U0001D10D")` returns "\U0