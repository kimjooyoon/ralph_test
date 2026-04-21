# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, reversing, and encoding
- **Data Analysis**: Provides numeric range generation and average calculation
- **Encoding/Decoding**: Implements base64, hex, and JSON serialization
- **AI Code Generation**: Produces AI-style code snippets from patterns

## Aggregate Boundaries
- `Split` operates on UTF-8 byte sequences (not Unicode code points)
- `Reverse` handles surrogate pairs and combining marks as single code points
- `EncodeBase64` converts byte slices to base64 strings
- `GenerateCode` produces code snippets from pattern strings

## Domain Events
- `Split` completes with UTF-8 byte splits
- `Reverse` completes with properly normalized Unicode
- `EncodeBase64` completes with valid base64 encoding
- `GenerateCode` completes with AI-style code snippet

## Invariants
- `Split("", "")` returns empty slice
- `Split(s, "")` for non-empty `s` returns byte splits
- `Reverse` preserves surrogate pairs as single code points
- `EncodeBase64` handles all valid UTF-8 input
<<END_FILE>> 

<<<