# Context Map

## Bounded Contexts
- **String Manipulation**: Handles splitting, reversing, joining, trimming, and replacing strings with UTF-8 byte/character awareness.
- **Unicode Processing**: Specializes in handling multi-byte characters (e.g., Chinese, emoji), surrogate pairs, and combining marks.
- **Pure Functions**: Ensures all helpers are stateless, side-effect-free, and testable via unit tests.

## Aggregate Boundaries
- `Split(s, sep)` operates on a string and separator to produce a slice of substrings, respecting UTF-8 byte/character boundaries.
- `Reverse(s)` transforms a string into its reverse, preserving surrogate pairs and combining marks.
- `Join(sep, parts)` combines a slice of strings into a single string with a separator, handling edge cases for empty inputs.

## Domain Events/Invariants
- **Split invariant**: For `sep == ""`, returns one string per UTF-8 byte; for non-empty `s`, returns one string per Unicode code point.
- **