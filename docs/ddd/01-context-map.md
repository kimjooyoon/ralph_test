# Context Map

## Bounded Contexts
- **String Manipulation**: Handles core string operations (Split, Reverse, Join, Trim, etc.)
  - Aggregate boundaries:
    - `Split` operates on string segments at UTF-8 byte level
    - `Reverse` handles Unicode code points with surrogate pair awareness
    - `Join` manages string concatenation with separator injection
  - Domain Events:
    - `StringSplitEvent` (when Split completes)
    - `StringReverseEvent` (when Reverse completes)
    - `StringJoinEvent` (when Join completes