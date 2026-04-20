# Context Map

## Bounded Contexts
- **String Manipulation**: Contains `Split`, `Join`, `Trim`, `Reverse`, `Replace`, `Repeat`, `Echo`  
- **Unicode Handling**: Specialized logic for UTF-8 byte splitting, surrogate pairs, and combining marks  
- **Numeric Operations**: `Range`, `ParseInt`, `Average`  
- **Encoding/Decoding**: Base64, Hex, JSON, YAML, AES encryption  
- **Pattern Matching**: Regex-like `Match`, `ContainsWildcard`  

## Aggregate Boundaries
- `Split` aggregates: UTF-8 byte splitting (Chinese, emoji)  
- `Reverse` aggregates: Surrogate pair handling  
- `Range` aggregates: Numeric range generation  
- `EncodeBase64` aggregates: Binary-to-string encoding  

## Domain Events/Invariants
- `Split` invariant: Empty separator → UTF-8 byte splitting  
- `Reverse` invariant: Surrogate pairs treated as single code points  
- `Range` invariant: Inclusive range generation  
- `ParseInt` invariant: Valid UTF-8 input required  

## Open Questions
- Should `Split` handle surrogate pairs as single code points or byte splits?  
- How to model `GenerateCode` and `Eval` as pure functions?  
