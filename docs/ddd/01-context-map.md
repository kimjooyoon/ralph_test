# Context Map

## Bounded Contexts
- **String Manipulation Utilities**: Handles core string operations (split, reverse, replace, etc.) with strict UTF-8 compliance.

## Aggregates
- `Split`: Splits strings by UTF-8 bytes (for multi-byte characters like Chinese)
- `Reverse`: Reverses strings while preserving surrogate pairs and combining marks
- `Replace`: Substitutes patterns in strings (pure function)
- `Match`: Basic regex-like pattern matching (pure function)

## Domain Events
- `StringSplitEvent`: Triggered when UTF-8 byte splitting completes
- `StringReverseEvent`: Triggered when reversal completes

## Invariants
- `Split` must treat empty string as empty slice
- `Reverse` must handle surrogate pairs as single code points
- All operations must use valid UTF-8 input

## Open Questions
- How to handle overlapping surrogate pairs in