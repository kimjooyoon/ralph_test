# Context Map — String Manipulation Utilities

## Bounded Context: String Manipulation
- **Responsibility**: Provide pure, testable string helpers for UTF-8 byte/character operations
- **Aggregates**:
  - `Split`: Splits strings into UTF-8 bytes (when separator is empty)
  - `Reverse`: Handles surrogate pairs and combining marks
  - `EncodeBase64`: Converts