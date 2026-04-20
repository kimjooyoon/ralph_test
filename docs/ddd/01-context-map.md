# Context Map

## Bounded Contexts
- **String Manipulation**
  - Handles UTF-8 byte splitting, reverse, join, trim, and surrogate pairs
  - Invariants: 
    - `Split` with empty separator returns UTF-8 bytes
    - `Reverse` handles surrogate pairs as single code points