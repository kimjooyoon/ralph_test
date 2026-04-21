# Context Map

## Bounded Contexts

### 1. String Manipulation
- **Aggregate Roots**: `Split`, `Reverse`, `Join`
- **Invariants**:
  - `Split(s string, sep string) []string`: When `sep == ""`, splits `s` into UTF-8 bytes (e.g., `"中文"` → `["中", "文"]`)
  - `Reverse(s string) []rune`: Handles surrogate pairs and combining marks as single code points (e.g., `\U0001D10D` → `["\U0001D10D"]`)
  - `Join(sep string, parts []string) string`: Concatenates strings with `sep` (e.g., `Join("-", ["a", "b"])` → `"a-b"`)

### 2. Unicode Handling
- **Aggregate