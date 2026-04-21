# Context Map

## Bounded Contexts
- **String Manipulation** (primary context)
  - Handles UTF-8 string operations: splitting, reversing, encoding, etc.
  - Aggregates:
    - `Split` (splits strings by UTF-8 bytes or separators)
    - `Reverse` (reverses strings with Unicode-aware handling)
    - `EncodeBase64` (encodes bytes to Base64)
    - `Range` (generates numeric ranges)
    - `Match` (pattern matching with