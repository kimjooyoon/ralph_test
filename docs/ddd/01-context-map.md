# Context Map

## Bounded Contexts
- **StringHelpers**: Core string manipulation functions (Split, Reverse, Join, Trim, Replace, etc.)
  - Aggregate: `StringOperation` (immutable, pure functions)
  - Invariants: 
    - `Split(s string, sep string) []string` must split by UTF-8 bytes for empty separator