# Context Map

## Bounded Contexts

### 1. String Manipulation
- **Description**: Core string operations with UTF-8 awareness
- **Aggregates**:
  - `StringOperations`: 
    - Invariant: `Split("中文", "")` returns ["中", "文"] (UTF-8 byte splitting)
    - Invariant: `Reverse("\U0001D10D")` handles surrogate pairs as single code points
    - Invariant: `ContainsWildcard("