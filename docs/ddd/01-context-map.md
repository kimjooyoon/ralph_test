# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, reversing, encoding, and pattern matching.  
- **Data Analysis**: Provides numeric range generation, averaging, and parsing.  
- **AI Code Generation**: Generates AI-style code snippets and evaluates expressions.  
- **Serialization/Encryption**: Offers JSON/YAML encoding, base64/hex conversion, and mock encryption.  

## Aggregate Boundaries
- **Split**: Splits strings into UTF-8 bytes (empty separator) or by pattern.  
- **Reverse**: Reverses strings while preserving surrogate pairs and combining marks.  
- **GenerateCode**: Produces AI-style code snippets