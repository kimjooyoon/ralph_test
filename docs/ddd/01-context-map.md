# Context Map — String Manipulation Helpers

## Bounded Context: String Manipulation Helpers
- **Primary responsibility**: Provide pure, testable string operations for domain experiments
- **Aggregates**:
  - `Split`: UTF-8 byte splitting for multi-byte characters (Chinese, emoji)
  - `Reverse`: Surrogate pair & combining mark handling
  - `EncodeBase64`: Binary-to-text encoding
  - `GenerateCode`: AI-style code snippet generation
  - `Average`: Numerical mean calculation
- **Invariants**:
  - `Split("中文", "")` → ["中", "文