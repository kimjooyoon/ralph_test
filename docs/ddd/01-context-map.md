# Context Map

## Bounded Contexts

### 1. String Manipulation
- **Aggregates**: `Split`, `Join`, `Trim`, `Reverse`
- **Invariants**:
  - `Split` must split UTF-8 bytes for multi-byte characters (e.g., Chinese, emoji)
  - `Reverse` must handle surrogate pairs and combining marks correctly
  - `Trim` must remove leading/trailing whitespace per Unicode definition
  - `Join` must concatenate strings with separator per UTF-8 rules

### 2. Unicode Handling
- **Aggreg