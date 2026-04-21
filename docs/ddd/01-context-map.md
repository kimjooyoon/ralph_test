# Context Map

## Bounded Contexts
1. **String Manipulation**
   - Responsible for UTF-8 byte/character splitting, reversing, and basic string transformations
   - Aggregates: `String` (with `Split`, `Reverse`, `Trim`, `Join`, `Repeat`, `Replace`)
   - Invariants: 
     - `Split` must split on UTF-8 bytes for multi-byte characters (e.g., Chinese, emoji)
     - `Reverse` must handle surrogate pairs and combining marks correctly
     - All string operations must preserve valid UTF-8 input

2. **Encoding/Decoding**
   - Handles base64, hex, and other encoding formats
   - Aggregates: `EncodedData` (with `EncodeBase64`, `EncodeHex`)
   - Invariants: 
     - Encoding/decoding must be pure functions with no I/O
     - Base64 encoding must handle non-ASCII input correctly

3. **AI Code Generation**
   - Generates AI-style code snippets from patterns
   - Aggregates: `CodeSnippet` (with `GenerateCode`)
   - Invariants: 
     - `GenerateCode` must return syntactically valid code
     - Must be a pure function with no I/O

## Context Boundaries
- `Split` boundary: UTF-8 byte vs Unicode