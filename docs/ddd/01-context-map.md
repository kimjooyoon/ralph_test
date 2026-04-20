# Context Map — String Manipulation Bounded Context

## Bounded Context
**String Manipulation** (domain layer)  
- Handles UTF-8 byte splitting, Unicode normalization, and surrogate pair processing  
- Pure functions with no I/O or external dependencies  

## Aggregate Boundaries
1. **Split Aggregate**  
   - `Split(s string, sep string)`  
   - Invariant: When `sep == ""`, splits by UTF-8 byte (not Unicode code point)  
   - Edge case: Empty string returns empty slice  

2. **Reverse Aggregate**  
   - `Reverse(s string)`  
   - Invariant: Handles surrogate pairs and combining marks as single code points  
   -