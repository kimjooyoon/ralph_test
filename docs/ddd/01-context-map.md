# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, reversing, and pattern matching (Split, Reverse, Match)
- **Unicode Handling**: Specializes in surrogate pairs, combining marks, and multi-byte character processing (Split, Reverse)
- **Numeric Operations**: Focuses on range generation and string-to-integer conversion (Range, ParseInt)
- **AI Code Generation**: Produces AI-style code snippets from patterns (GenerateCode)

## Aggregate Boundaries
- **String Manipulation Aggregate**: 
  - `Split(s, sep)` - UTF-8 byte splitting for multi-byte characters
  - `Reverse(s)` - Surrogate pair-aware reversal
  - `Match(pattern, s)` - Regex-like pattern matching
- **Numeric Aggregate**:
  - `Range(start, end)` - Inclusive numeric range generation
  - `ParseInt(s)` - String-to-integer conversion with validation

## Domain Events
- `StringSplitEvent` - Triggered when UTF-8 byte splitting