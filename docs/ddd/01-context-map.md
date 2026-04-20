# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, reverse, and surrogate pair handling
- **Unicode Handling**: Specializes in multi-byte character operations (Chinese, emoji)
- **Encoding/Decoding**: Base64, hex, and other data transformations

## Aggregate Boundaries
- `Split` aggregates: UTF-8 byte splitting with empty separator
- `Reverse` aggregates: Surrogate pair preservation
- `GenerateCode` aggregates: AI-style code snippet generation

## Domain Events
- `CharacterSplitEvent`: Triggered when UTF-8 bytes are split
- `SurrogatePreservedEvent`: Triggered when surrogate pairs are maintained
- `CodeGeneratedEvent`: Triggered when AI-style code is produced

## Invariants
- `SplitInvariant`: Must split by UTF-8 bytes, not code points
- `ReverseInvariant`: Must preserve surrogate pairs as single code points
- `CodeGenInvariant`: Must produce valid code snippets without I/O

## Open Questions
- How to handle invalid UTF-8 sequences