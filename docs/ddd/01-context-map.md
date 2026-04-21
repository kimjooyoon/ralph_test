# Context Map

## Bounded Contexts
1. **String Manipulation**
   - Aggregates: `Split`, `Reverse`, `Join`, `Trim`, `Repeat`, `Replace`
   - Invariants: 
     - `Split` must split UTF-8 bytes for multi-byte characters (e.g., Chinese, emoji)
     - `Reverse` must handle surrogate pairs and combining marks correctly
     - All operations must preserve