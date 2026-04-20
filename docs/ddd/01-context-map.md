# Context Map

## Bounded Contexts

1. **String Manipulation**  
   - Handles UTF-8, Unicode, and byte-level operations  
   - Includes: `Split`, `Reverse`, `Join`, `Trim`, `Repeat`, `Replace`  
   - Aggregates: `StringSegment`, `UnicodeCodePoint`, `ByteRange`

2. **Numeric Operations**  
   - Pure functions for range generation and numeric analysis  
   - Includes: `Range`, `Average`  
   - Aggregates: `NumericInterval`, `MeanValue`

3. **Encoding/Decoding**  
   - Base64, Hex, and mock image/JSON/YAML serialization  
   - Includes: `EncodeBase64`, `EncodeHex`, `DecodeImage`, `SerializeJSON`, `SerializeYAML`  
   - Aggregates: `EncodedData`, `BinaryPayload`

4. **AI-Generated Code**  
   - Mock functions for code generation and evaluation  
   - Includes: `GenerateCode`, `Eval`, `DecodeImage`  
   - Aggregates: `CodeSnippet`, `ExecutionResult`

## Aggregate Boundaries

- `Split` operates on UTF-8 byte sequences, treating each byte as a segment  
- `Reverse` handles Unicode code points, preserving surrogate pairs  
- `Range` generates numeric intervals as pure functions  
- `EncodeBase64` converts binary data to ASCII strings  

## Domain Events

- `StringSegmentCreated`: When `Split`