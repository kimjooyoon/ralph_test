# Context Map

## Bounded Context: String Manipulation Utilities
- **Primary responsibility**: Provide pure, testable string operations with explicit UTF-8 handling
- **Key aggregates**:
  - `Split` (splits strings by UTF-8 bytes or separators)
  - `Reverse` (handles surrogate