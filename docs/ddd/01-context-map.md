# Context Map

## Bounded Contexts

### 1. String Manipulation
- **Responsibilities**: UTF-8 byte splitting, Unicode handling, surrogate pairs, reversing, trimming, joining, repeating, and basic pattern matching.
- **Aggregate Roots**: 
  - `Split` (splits strings by UTF-8 bytes or code points)
  - `Reverse` (handles surrogate pairs and combining marks)
  - `Trim` (removes whitespace from strings)
  - `Join` (concatenates strings with separators)