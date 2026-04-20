# Context Map — String Manipulation Domain

## Bounded Context: String Manipulation
- **Primary responsibility**: Provide pure, testable string operations with UTF-8 awareness
- **Aggregates**:
  - `Split` (splits strings by UTF-8 bytes/code points)
  - `Reverse` (handles surrogate pairs, combining marks)
  - `GenerateCode` (AI-style code generation)
  - `Average` (numeric operations)
  - `DecodeImage` (simulated