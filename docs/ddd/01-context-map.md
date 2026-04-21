# Context Map

## Bounded Contexts

### 1. String Manipulation Helpers
- **Responsibility**: Provide pure, testable string operations for domain experiments
- **Aggregates**:
  - `Split` (UTF-8 byte splitting for multi-byte characters)
  - `Reverse` (surrogate pair & combining mark handling)
  - `GenerateCode` (AI-style code snippet generation)
- **Invariants**:
  - `Split` must split UTF-8 bytes, not Unicode code points
  - `Reverse` must preserve surrogate pairs as single code points
  - `GenerateCode` must return syntactically valid code snippets

### 2. Numeric Operations
- **Responsibility**: Handle numeric