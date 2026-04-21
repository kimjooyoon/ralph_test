# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, surrogate pairs, and Unicode code points.  
  - **Aggregate Roots**: `Split`, `Reverse`, `EncodeBase64`  
  - **Domain Events**: `CharacterSplit`, `SurrogatePairProcessed`, `Base64Encoded`  
  - **Invariants**:  
    - `Split(s, "")` returns one string per UTF-8 byte (not Unicode code point)  
    - `Reverse` correctly handles surrogate pairs and combining marks  
    - `EncodeBase64` produces valid Base64 for UTF-