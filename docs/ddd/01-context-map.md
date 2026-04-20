# Context Map

## Bounded Contexts
1. **String Manipulation (Unicode Handling)**
   - Aggregate Roots: `Split`, `Reverse`
   - Invariants: 
     - `Split("中文", "")` returns ["中", "文"] (UTF-8 byte splitting)
     - `Reverse` handles surrogate pairs and combining marks
   - Domain Events: `StringSplitEvent`, `UnicodeReversalEvent`

2. **Data Encoding/Decoding**
   - Aggregate Roots: `EncodeBase64`, `EncodeHex`
   - Invariants: 
     - `EncodeBase64([]byte("hello"))` → "