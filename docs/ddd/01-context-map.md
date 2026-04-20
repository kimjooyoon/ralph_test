# Context Map

## Bounded Contexts

### String Manipulation
- **Responsibility**: Core string operations (split, join, reverse, replace, trim)
- **Aggregates**:
  - `String` (immutable value object)
  - `Splitter` (handles UTF-8 byte/character splitting)
  - `Reverser` (handles surrogate pairs, combining marks)
- **Domain Events**:
  - `StringSplitEvent` (when split operation completes)
  - `StringReversedEvent` (when reversal completes)
- **Invariants**:
  - `SplitMustNotModifyInput` (split returns new strings)
  -