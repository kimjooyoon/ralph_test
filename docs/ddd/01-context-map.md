# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, surrogate pairs, and Unicode edge cases.  
- **Code Generation**: Produces AI-style code snippets via pure functions.  
- **Numeric Operations**: Computes averages and generates ranges.  
- **Encoding/Decoding**: Base64, hex, and image decoding helpers.  

## Aggregate Boundaries
- **Split**: Splits strings by UTF-8 bytes (not code points).  
- **Reverse**: Handles surrogate pairs and combining marks as single units.  
- **GenerateCode**: Pure function with