# Context Map

## Bounded Context: String Manipulation
- **Primary responsibility**: Provide pure, testable string utilities with UTF-8 awareness
- **Key aggregates**:
  - `Split` (splits by UTF-8 bytes, not code points)
  - `Reverse` (handles surrogate pairs and combining marks)
  - `EncodeBase64` (pure encoding without I/O)
  - `Match` (regex-like pattern matching)
  - `ContainsWildcard` (wild