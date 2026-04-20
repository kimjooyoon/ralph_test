# Context Map

## Bounded Contexts

1. **String Manipulation**
   - Aggregates: `Split`, `Reverse`, `Join`, `Trim`, `Repeat`, `Replace`
   - Invariants: 
     - `Split` must split UTF-8 bytes, not Unicode code points
     - `Reverse` must treat surrogate pairs as single code points