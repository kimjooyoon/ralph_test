# Context Map

## Bounded Contexts

1. **DSL Helpers** (primary context)
   - Aggregate: `StringManipulation` (root)
   - Subcontexts: 
     - `UnicodeHandling` (for Split/Reverse)
     - `NumericOperations` (for Range/ParseInt)