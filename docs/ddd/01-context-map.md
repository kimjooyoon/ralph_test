# Context Map

## Bounded Contexts
- **String Manipulation** (primary context for string helpers)
- **Data Encoding** (for base64, hex, etc.)
- **Numeric Operations** (range, averages)
- **Pattern Matching** (regex-like)
- **Serialization** (JSON, YAML)
- **Logging** (mock logging)

## Aggregates & Boundaries
- **String Manipulation Aggregate**
  - `Split(s, sep)` - Splits strings by UTF-8 bytes (when sep is empty)
  - `Reverse(s)` - Handles surrogate pairs and combining marks
  - `Join(sep, parts)` - Combines strings with separator
  - `Trim(s)` - Removes whitespace (UTF-8 aware)
  - `Repeat(s, n)` - Reproduces string n times
  - `Replace(s, old, new)` - Replaces substrings

- **Data Encoding Aggregate