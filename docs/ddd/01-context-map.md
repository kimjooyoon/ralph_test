# Context Map

## Bounded Contexts
- **StringHelpers**: Handles string splitting, reversing, and encoding
  - Aggregate: `Split` (UTF-8 byte splitting for multi-byte characters)
  - Aggregate: `Reverse` (surrogate pair and combining mark handling)
  - Aggregate: `EncodeBase64` (pure encoding function)
- **NumericOps**: Manages numeric range generation and averaging
  - Aggregate: `Range` (inclusive numeric range generation)
  - Aggregate: `Average` (floor division logic)
- **CodexAI**: Generates AI-style code snippets
  - Aggregate: `GenerateCode` (pattern-based code generation)
- **ImageDecoding**: Simulates image decoding from string data
  - Aggregate: `DecodeImage` (PNG header extraction)

## Domain Events
- `StringSplitCompleted`: Triggered when string is split into byte segments
- `RangeGenerated`: When numeric range is successfully generated
- `CodeGenerated`: