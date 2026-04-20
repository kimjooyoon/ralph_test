# Context Map

## Bounded Contexts
1. **String Manipulation**
   - Aggregates: `Split`, `Join`, `Trim`, `Reverse`, `Replace`, `Repeat`
   - Boundaries: 
     - `Split` handles UTF-8 byte splitting (not Unicode code points)
     - `Reverse` must process surrogate pairs and combining marks
     - `Match`/`ContainsWildcard` operate on regex-like pattern matching
   - Domain Events: 
     - `StringSplitEvent` (when Split completes)
     - `PatternMatchedEvent` (when regex-like pattern is validated)

2. **Numeric Operations**
   - Aggregates: `Range`, `Average`
   - Boundaries: 
     - `Range` generates inclusive integer sequences
     - `Average` computes floor of mean for integer slices

3. **Encoding/Decoding**
   - Aggregates: `EncodeBase64`, `DecodeImage`
   - Boundaries: 
     - `EncodeBase64` produces valid Base64 strings
     - `DecodeImage` simulates image decoding from byte slices

## Ubiquitous Language
- **Split**: "Split the string by UTF-8 bytes" (not code points)
- **Reverse**: "Reverse the string while preserving surrogate pairs"
- **Match**: "Check if string matches regex-like pattern"
- **ContainsWildcard**: "Check if string contains wildcard pattern (*?)