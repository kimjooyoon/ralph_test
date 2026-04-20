# Context Map

## Bounded Contexts
- **String Manipulation**: Contains `Split`, `Join`, `Trim`, `Reverse`, `Replace`, `Repeat`, `Echo` for UTF-8 byte/character operations
- **Unicode Processing**: Specializes in surrogate pairs, combining marks, and multi-byte character handling
- **Data Transformation**: Includes numeric helpers (`Range`, `Average`, `ParseInt`), encoding/decoding (`Base64`, `Hex`), and serialization (`JSON`, `YAML`)
- **AI Code Generation**: Mocks code evaluation and generation via `Eval`, `GenerateCode`, and `DecodeImage`

## Aggregate Boundaries
- `Split` aggregates: UTF-8 byte splitting with empty separator
- `Reverse` aggregates: Unicode code point reversal with surrogate pair handling
- `Join` aggregates: String concatenation with delimiter injection
- `Trim` aggregates: Leading/trailing character removal

## Domain Events
- `StringSplitEvent`: Triggered when `Split` completes byte/character separation
- `UnicodeReversalEvent`: Fired after `Reverse` processes surrogate pairs
- `CodeGenerationEvent`: Occurs when `GenerateCode` produces AI-style snippets

## Invariants
- `Split(s, "")` must return one element per UTF-8 byte
- `Reverse` must treat surrogate pairs as single code points
- `Join` must preserve delimiter integrity between elements
- `Trim` must remove only matching leading/trailing characters

## Open Questions
- How