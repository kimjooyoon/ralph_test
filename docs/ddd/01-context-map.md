# Context Map

## Bounded Contexts

1. **String Manipulation**
   - **Ubiquitous Language**: Split, Reverse, Join, Trim, Replace
   - **Aggregate Boundaries**: 
     - `Split(s, sep)` handles UTF-8 byte splitting (ASCII) vs code point splitting
     -