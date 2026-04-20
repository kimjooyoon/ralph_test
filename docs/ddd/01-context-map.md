# Context Map

## Bounded Contexts
- **String Manipulation**: Handles core string operations (split, reverse, trim, etc.)
- **Numeric Processing**: Focuses on number-based operations (average, range, parsing)
- **Unicode Handling**: Specializes in multi-byte character and surrogate pair processing
- **Code Generation**: Implements AI-style code snippet generation (GenerateCode)
- **Image Simulation**: Provides mock image decoding (DecodeImage)

## Aggregate Boundaries
- **String Manipulation**:
  - `Split(s, sep)` splits strings by UTF-8 bytes (not code points)
  - `Reverse(s)` handles surrogate pairs and combining marks
  - `Trim(s)` removes whitespace from both ends
- **Numeric Processing**:
  - `Average(nums)` computes mean with floor division
  - `Range(start, end)` generates inclusive integer ranges
  - `ParseInt(s)` converts strings to integers with validation
- **Unicode Handling**:
  - `Split` must correctly process UTF-8 byte sequences for Chinese, emoji, etc.
  - `Reverse` must treat surrogate