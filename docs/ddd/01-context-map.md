# Context Map

## Bounded Contexts
- **String Manipulation** (Split, Reverse, Trim, Join, Replace, Repeat)
  - Aggregate roots: `Split`, `Reverse`
  - Invariants: 
    - `Split("中文", "")` returns `["中", "文"]` (UTF-8 byte splitting)
    - `Reverse("\U0001D10D")` returns same emoji (surrogate pair handling)
- **Data Encoding** (Base64, Hex, Encode/Decode)
  - Aggregate roots: `EncodeBase64`, `EncodeHex`