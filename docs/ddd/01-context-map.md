# Context Map — String Manipulation Domain

## Bounded Contexts
- **String Manipulation**: Core domain for UTF-8 byte/character operations
- **Unicode Handling**: Specialized context for surrogate pairs, combining marks

## Aggregates
- **String Splitter**: Splits strings by UTF-8 bytes (when separator is empty)
  - Invariant: Must not split multi-byte characters (e.g., Chinese)
  - Event: `StringSplitEvent` when split operation completes
- **Unicode Validator**: Ensures valid UTF-8 sequences
  - Invariant: All input strings must be valid UTF-8
  - Event: