# Context Map

## Bounded Contexts
1. **String Manipulation**  
   - Primary: `Split`, `Reverse`, `EncodeBase64`, `Match`, `Replace`, `Join`, `Repeat`, `Trim`, `Range`, `ParseInt`, `Average`, `DecodeImage`, `Eval`, `GenerateCode`
   - Secondary: N/A

2. **Unicode Handling**  
   - Primary: `Split`, `Reverse`, `Match`, `EncodeBase64`, `DecodeImage`
   - Secondary: N/A

## Aggregates & Invariants
- **String Manipulation Aggregate**  
  - Invariant: All operations must preserve UTF-8 validity (no invalid byte sequences)
  - Invariant: `Split` with empty separator must split by UTF-8 bytes, not Unicode code points
  - Invariant: `Reverse` must treat surrogate pairs as single code points

- **Unicode Aggregate**  
  - Invariant: Surrogate pairs must be treated as single code points in all operations
  - Invariant: Emoji and combining marks must be preserved as single code points

## Domain Events
- `StringSplitEvent`: Triggered when `Split` is called with empty separator
- `UnicodeReverseEvent`: Triggered when `Reverse` is called on UTF-8 strings
- `Base64EncodeEvent`: Triggered when `EncodeBase64` is called
- `ImageDecodeEvent`: Triggered when `DecodeImage` is called

## Open Questions
- How to handle invalid UTF-8 input in `Split` and `Reverse