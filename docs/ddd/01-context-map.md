# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, reversing, and pattern matching (bounded by `Split`, `Reverse`, `Match` functions)
- **Unicode Handling**: Specializes in surrogate pairs, combining marks, and multi-byte character processing (bounded by `Split`, `Reverse` invariants)
- **Data Encoding**: Focuses on base6