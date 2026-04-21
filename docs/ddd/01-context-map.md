# Context Map

## Bounded Contexts

### 1. String Manipulation
- **Description**: Core string operations with UTF-8 awareness
- **Aggregate Roots**:
  - `Split` (UTF-8 byte splitting)
  - `Reverse` (surrogate pair handling)
  - `Join` (string