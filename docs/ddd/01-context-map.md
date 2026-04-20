# Context Map

## Bounded Contexts
- **String Manipulation**: Handles core string operations (Split, Reverse, Replace, Trim, Join, etc.)
- **Pattern Matching**: Regex-like matching (Match function)
- **Unicode Handling**: Specialized for multi-byte characters, surrogate pairs, and combining marks
- **Numeric Utilities**: Range generation, parsing, and basic math operations

## Aggregate Boundaries
- **String Manipulation**:
  - `Split` (UTF-8 byte splitting for multi-byte characters)
  - `Reverse` (surrogate pair and combining mark handling)
  - `Replace` (pattern-based replacement)
  - `Trim` (leading/trailing whitespace