# Context Map — DSL Helpers

## Bounded Contexts
1. **String Manipulation** (primary context)
   - Functions: Split, Join, Trim, Reverse, Replace, Repeat, Match, ContainsWildcard
   - Invariants: 
     - `Split` handles UTF-8 bytes for multi-byte characters (e.g., Chinese, emoji)
     - `Reverse` respects surrogate pairs and combining marks
     - All functions are pure (no I/O, no clocks, no network)

2. **Numeric Helpers** (sub-context of String Manipulation)
   - Functions: Range, ParseInt
   - Invariants: 
     - `Range` generates inclusive numeric sequences
     -