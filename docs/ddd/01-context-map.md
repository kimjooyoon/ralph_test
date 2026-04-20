# Context Map

## Bounded Contexts
1. **String Manipulation**
   - Aggregate: `Split`
   - Invariants: 
     - `Split(s, "")` returns UTF-8 byte splits (e.g., "中文" → ["中", "文"])
     - `Split` handles surrogate pairs as single code points
   - Domain Events: 
     - `StringSplitEvent` (when split operation completes)

2. **Code Generation**
   - Aggregate: `GenerateCode`
   - Invariants: 
     - Returns AI-style code snippets via pattern matching
     - Pure function with no I/O
   - Domain Events: