# Context Map

## Bounded Contexts

1. **String Manipulation**
   - Functions: Split, Reverse, Join, Trim, Repeat, Replace
   - Invariants: 
     - `Split` must split UTF-8 bytes for multi-byte characters (e.g., Chinese, emoji)
     - `Reverse` must handle surrogate pairs and combining marks
     - `Join` must preserve UTF-8 byte boundaries
     - `Trim` must remove whitespace from UTF-8 strings
     - `Repeat` must handle numeric ranges and string repetition
     - `Replace` must support pattern-based string substitution

2. **Unicode Handling**
   - Functions: Split (with byte