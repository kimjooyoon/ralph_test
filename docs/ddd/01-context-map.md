# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, reversing, and surrogate pair processing (Split, Reverse, TestSplitSurrogatePairs)
- **AI Code Generation**: Generates code snippets via pattern matching (GenerateCode)
- **Data Analysis**: Computes statistical averages (Average)
- **Unicode Processing**: Specialized handling of multi-byte characters and edge cases (Split, Reverse, TestSplitSurrogatePairs)

## Aggregate Boundaries
- **String Manipulation Aggregate**: 
  - `Split(s string, sep string) []string` 
  - `Reverse(s string) string` 
  - `TestSplitSurrogatePairs()` (ensures surrogate pairs are treated as single code points)
- **Code Generation Aggregate**: 
  - `GenerateCode(pattern string) string` 
- **Data Analysis Aggregate**: 
  - `Average(nums []int) int`

## Domain Events
- **String Split Event**: Triggered when UTF-8 byte splitting occurs
-