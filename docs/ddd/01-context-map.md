# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, surrogate pairs, and Unicode normalization. Exposes `Split`, `Reverse`, `EncodeBase64`, and `ParseInt` as pure functions with no side effects.

## Aggregate Boundaries
- **String Splitting**: `Split(s string, sep string)` splits UTF-8 bytes (not Unicode code points) when `sep == ""`. Must preserve multi-byte characters (e.g., Chinese, emoji) as single elements.
- **Unicode Processing**: `Reverse(s string)` must handle surrogate pairs and combining marks as single code points. Invalid UTF-8 input is treated as invalid sequences (per specs).

## Ubiquitous Language