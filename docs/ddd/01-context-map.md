# Context Map

## Bounded Contexts

### 1. String Processing
- **Ubiquitous Language**: Split, Reverse, Trim, Join, Replace, Repeat, EncodeBase64, ParseInt, EncodeHex, EncodeJSON, EncodeYAML
- **Aggregates**:
  - `Split` (splits strings by UTF-8 bytes/code points)
  - `Reverse` (handles surrogate pairs, combining marks)
  - `Trim` (removes whitespace)
  - `Join` (concatenates strings with separator)
  - `Replace` (substitutes substrings)
  - `Repeat` (repeats strings)
- **Invariants**:
  - `Split("", "")` → `[]string{""}`
  - `Split("中文", "")` → `["中", "文"]` (UTF-8 byte split)
  - `Reverse("\U0001D10D")` → `"\U0001D10D"` (surrogate pair treated