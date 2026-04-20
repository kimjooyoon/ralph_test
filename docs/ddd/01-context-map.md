# Context Map — String Manipulation Helpers

## Bounded Context: String Manipulation DSL
- **Primary responsibility**: Provide pure, testable string helpers for domain experiments
- **Key aggregates**:
  - `Split` (splits strings by UTF-8 bytes/characters)
  - `Reverse` (handles surrogate pairs and combining marks)
  - `EncodeBase64` (encodes bytes to base64)
  - `Range` (generates numeric ranges)
- **Domain events**:
  - `StringSplitEvent` (when Split completes)
  - `StringReverseEvent` (when Reverse completes)
- **Invariants**:
  - `Split(s, "")` must split by UTF-8 bytes (not code points)
  - `Reverse` must preserve surrogate pairs as single code points
  - All helpers must be pure functions (no I/O, no side effects)

## Context Boundaries
- **Internal**: All string helpers are internal to this bounded context
- **External**: No external dependencies (pure functions only)
<<<END