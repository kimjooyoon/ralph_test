# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 splitting, reversing, and pattern matching (includes `Split`, `Reverse`, `Match`)
- **Numeric Operations**: Range generation, averaging, and parsing (`Range`, `Average`, `ParseInt`)
- **Encoding/Decoding**: Base64, hex, and image decoding (`EncodeBase64`, `DecodeImage`)
- **Unicode Handling**: Surrogate pairs, combining marks, and multi-byte character support (`Split`, `Reverse`)
- **AI-Generated Code**: Mock code generation and evaluation (`GenerateCode`, `Eval`)

## Aggregate Boundaries
- `Split` aggregates UTF-8 byte splitting logic
- `Reverse` aggregates Unicode-aware reversal
- `Average` aggregates numeric mean calculation
- `DecodeImage` aggregates image data simulation

## Domain Events
- `StringSplitEvent` (triggered by `Split`)
- `UnicodeReversedEvent` (triggered by `Reverse`)
- `RangeGeneratedEvent` (triggered by `Range`)
- `ImageDecodedEvent` (triggered by `DecodeImage`)

## Invariants
- `Split` must split UTF-8