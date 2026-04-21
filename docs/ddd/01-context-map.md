# Context Map

## Bounded Contexts

1. **String Manipulation**
   - Aggregates: `Split`, `Join`, `Trim`, `Reverse`
   - Invariants: 
     - `Split` with empty separator splits on UTF-8 bytes (not Unicode code points)
     - `Reverse` handles surrogate pairs and combining marks
     - `Trim` removes whitespace per Unicode definition

2. **Unicode Handling**
   - Aggregates: `Split`, `Reverse`, `Range`
   - Invariants: 
     - Correctly processes