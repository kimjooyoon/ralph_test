# Context Map

## Bounded Contexts
- **String Manipulation** (primary context): Handles UTF-8 string operations, including splitting, reversing, encoding, and pattern matching. This context contains all string-related utilities.

## Aggregate Boundaries
- **String Splitter**: Splits strings by UTF-8 bytes or code points. Aggregates are immutable and stateless.
- **Unicode Reverser**: Reverses strings while preserving surrogate pairs and combining marks. Aggregates are immutable and stateless.
- **Encoding Utilities**: Base64, hex, and other encoding/decoding operations. Aggregates are immutable and stateless.

## Domain Events
- `StringSplitEvent`: Triggered when a string is split by UTF-8 bytes or code points.
- `UnicodeReverseEvent`: Triggered when a string is reversed with proper Unicode handling.
- `EncodeEvent`: Triggered when data is encoded into a string format.

## Ubiquitous Language
- **UTF-8 byte splitting**: Splitting by individual UTF-8 bytes (not code points).
- **Surrogate pairs**: Pairs of Unicode code points used to represent characters outside the Basic Multilingual Plane.
- **Combining marks**: Diacritical marks that modify base characters (e.g., accents).
- **Code points**: Individual Unicode characters, including surrogate pairs.

## Open Questions
- Should `Split` with empty separator be a separate bounded context or part of the main string manipulation context?
- How