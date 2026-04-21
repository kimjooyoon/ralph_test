# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, surrogate pairs, and Unicode normalization. (Includes `Split`, `Reverse`, `EncodeBase64`)
- **Data Transformation**: Pure functions for encoding/decoding, serialization, and numeric operations. (Includes `EncodeBase64`, `Range`, `ParseInt`)
- **Pattern Matching**: Regex-like checks and wildcard substring searches. (Includes `Match`, `ContainsWildcard`)

## Aggregates
- **String Segmenter**: Ensures UTF-8 byte splitting for multi-byte characters (e.g., Chinese, emoji). (Invariant: `Split(s, "")` returns one element per UTF-8 byte)
- **Unicode Validator**: Ensures surrogate pairs and combining marks are treated as single code points. (Invariant: `Reverse("\U0001D10D")` preserves surrogate pairs)
- **Encoding Context**: Ensures base64 and hex encoding/decoding follows UTF-8 byte boundaries. (Invariant: `EncodeBase64([]byte("hello"))` produces "aGVsbG8=")

## Domain Events
- `StringSplitEvent`: Triggered when `Split` completes with UTF-8 byte splitting.
- `UnicodeReverseEvent`: Triggered when `Reverse` completes with surrogate pair preservation.
- `EncodeCompleteEvent`: Triggered when encoding