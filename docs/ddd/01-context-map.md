# Context Map

## Bounded Contexts
- **String Manipulation**: Handles core string operations (split, join, reverse, trim)
- **Pattern Matching**: Implements regex-like pattern matching (Match, ContainsWildcard)
- **Unicode Processing**: Specialized handling for Unicode (surrogate pairs, multi-byte chars)
- **Data Transformation**: Includes numeric helpers (Range, Average) and encoding/decoding

## Aggregates
- **String**: Core domain object for all manipulation operations
- **Pattern**: Represents regex-like patterns for matching
- **UnicodeCodePoint**: Represents Unicode code points for proper splitting
- **NumericRange**: Represents ranges of integers for data analysis

## Domain Events
- `StringSplit` (when Split operation completes)
- `PatternMatch` (when a