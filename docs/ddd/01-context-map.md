# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, Unicode reversal, and range generation for domain strings.
- **Code Generation**: Produces mock AI code snippets via pattern matching.
- **Encoding/Decoding**: Base64, hex, and other data transformations.

## Aggregate Boundaries
- **Split**: Splits strings by UTF-8 bytes (empty separator) or specific patterns. Invariant: returns byte-level segments for multi-byte characters.
- **Reverse**: Reverses strings while preserving surrogate pairs and combining marks. Invariant: treats surrogate pairs as single code points.
- **Range**: Generates inclusive integer ranges. Invariant: returns sequential integers between start and end.

## Domain Events
- `StringSplitEvent`: Triggered when a string is split by UTF-8 bytes.
- `CodeGeneratedEvent`: Triggered when AI-style code is generated via pattern matching.
- `RangeGeneratedEvent`: Triggered when an integer range is successfully generated.

## Open Questions
- How