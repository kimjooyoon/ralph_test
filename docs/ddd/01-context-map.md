# Context Map

## Bounded Contexts

### 1. String Manipulation
- **Key Functions**: Split, Reverse, Trim, Join, Replace, Repeat
- **Invariants**:
  - `Split(s, sep)` splits by UTF-8 byte for empty separator (e.g., "中文" → ["中", "文"])
  - `Reverse(s)` handles surrogate pairs and