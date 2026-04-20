# Context Map

## Bounded Contexts
1. **String Manipulation**
   - Functions: Split, Reverse, Replace, Trim, Join, Repeat, Echo
   - Invariants: 
     - `Split` handles UTF-8 bytes for multi-byte characters (e.g., Chinese, emoji)
     - `Reverse` correctly processes surrogate pairs and combining marks
     - All string operations are pure (no I/O, no clocks, no network)

2. **Regex-L