# Context Map

## Bounded Contexts
- **String Manipulation**: Handles splitting, reversing, trimming, joining, and replacing strings. 
  - `Split` splits strings by UTF-8 bytes (for empty separator).
  - `Reverse` reverses strings while preserving surrogate pairs and combining marks.
  - `Trim` removes leading/trailing whitespace.
  - `Join` concatenates strings with a separator.