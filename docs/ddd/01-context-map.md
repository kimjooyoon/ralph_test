# Context Map

## Bounded Contexts
1. **String Manipulation**
   - Aggregates: `Split`, `Join`, `Reverse`, `Trim`, `Replace`
   - Invariants: 
     - `Split("中文", "")` returns ["中", "文"] (UTF-8 byte splitting)
     - `Reverse("\U0001D10D")` treats surrogate pairs as single code points
     - `Trim("  abc