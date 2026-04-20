# Context Map — DSL Helpers

## Bounded Context: DSL Helpers
- **Primary responsibility**: Provide pure, testable string manipulation functions for domain experiments.
- **Key aggregates**:
  - `Split`: Splits strings by UTF-8 bytes (for multi-byte characters like Chinese)
  - `Reverse`: Handles surrogate pairs and combining marks
  - `Match`: Regex-like pattern matching with explicit syntax
  - `Replace