# Context Map

## Bounded Contexts
- **StringHelpers**: Provides pure string manipulation functions for UTF-8 and Unicode handling

## Aggregates
- **StringSplitter**: Splits strings by UTF-8 bytes or separators
  - Invariant: `Split("中文", "")` returns `["中", "文"]`
- **StringReverser**: Reverses strings with proper Unicode handling
  - Invariant: `Reverse("\U0001D10D")` returns `"\U0001D10D"` (surrogate pairs treated as single code points)

## Domain Events
- `StringSplitEvent`: Triggered when a string is split by UTF-8 bytes
- `StringReverseEvent`: Triggered when a string is reversed with Unicode awareness

## Ubiquitous Language
- **