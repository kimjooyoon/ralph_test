# Context Map

## Bounded Contexts
1. **String Manipulation**
   - Aggregates: StringSplitter, StringReverser, CodeGenerator
   - Domain Events: `StringSplitComplete`, `CodeGenerated`
   - Invariants:
     - `Split(s, "")` must return one string per UTF-8 byte (ASCII: 1 byte/char, Chinese: 3 bytes/char)
     - `Reverse` must handle surrogate pairs as single code points
     - `GenerateCode` must produce valid syntax without I/O

2. **Unicode Handling**
   - Aggregates: UnicodeValidator, SurrogatePairHandler
   - Domain Events: `UnicodeValidationFailed`, `SurrogatePairProcessed`
   - Invariants:
     - All UTF-8 inputs must be valid sequences
     - Surrogate pairs must be treated as single code points
     - Multi-byte characters must not split across bytes

3. **Data Transformation**
   - Aggregates: DataEncoder, DataDecoder
   - Domain Events: `DataEncoded`, `DataDecoded`
   - Invariants:
     - Base