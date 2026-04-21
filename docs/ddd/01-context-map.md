# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, surrogate pairs, and Unicode edge cases.  
  - Aggregates: `Split`, `Reverse`, `Trim`  
  - Invariants:  
    - `Split(s, sep)` returns UTF-8 bytes for `sep == ""`  
    - `Reverse(s)` preserves surrogate pairs as single code points  
    - `Trim(s)` removes whitespace per Unicode definition  

- **Code Generation**: Produces AI-style code snippets.  
  - Aggregates: `GenerateCode`  
  - Invariants:  
    - Output is pure function