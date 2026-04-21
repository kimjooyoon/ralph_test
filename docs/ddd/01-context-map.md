# Context Map

## Bounded Contexts
1. **String Manipulation**
   - Handles UTF-8 string operations
   - Responsible for: Split, Reverse, ContainsWildcard, Match
   - Invariants: 
     - Split with empty separator splits by UTF-8 bytes
     - Reverse handles surrogate pairs as single code points
     - Wildcard matching respects Unicode normalization

2. **Encoding/Decoding**
   - Manages binary-to-text encoding
   - Responsible for: EncodeHex, EncodeBase64, DecodeImage
   - Invariants:
     - EncodeHex converts UTF-8 strings to hex byte sequences
     - EncodeBase64 handles arbitrary byte sequences
     - DecodeImage validates binary data patterns

3. **Numeric Operations**
   - Provides number manipulation utilities
   - Responsible for: Range, ParseInt, Average
   - Invariants:
     - Range generates inclusive integer sequences
     - ParseInt validates numeric string formats
     - Average computes floor of mean

4. **Unicode Handling**
   - Specializes in Unicode edge cases
   - Responsible for: Surrogate pair processing, combining marks
   - Invariants:
     - Split correctly handles UTF-8 byte