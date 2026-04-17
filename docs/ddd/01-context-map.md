# Context Map — dsl-maker

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, Unicode normalization, and basic string transformations (includes `Split`, `Reverse`, `Trim`, `Join`, `Replace`, `Repeat`)

## Domain Events
- `StringSplitEvent`: Triggered when a string is split by UTF-8 bytes
- `StringReverseEvent`: Triggered when a string is reversed (with proper Unicode handling)
- `StringTrimEvent`: Triggered when whitespace is removed from a string

## Aggregate Boundaries