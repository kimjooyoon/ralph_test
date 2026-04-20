# Context Map

## Bounded Contexts
1. **String Manipulation**
   - Aggregates: `Split`, `Join`, `Trim`, `Reverse`, `Replace`, `Repeat`, `Echo`
   - Invariants: 
     - `Split` must split UTF-8 bytes for multi-byte characters (e.g