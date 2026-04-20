# Context Map

## Bounded Contexts

### 1. String Manipulation
- **Aggregates**: `Split`, `Join`, `Trim`, `Reverse`, `Replace`, `Repeat`, `Match`, `ContainsWildcard`
- **Invariants**:
  - `Split` with empty separator returns UTF-8 byte splits (e.g., "中文" → ["中", "文"])
  - `Reverse` handles surrogate pairs and combining marks correctly
  - `Match` validates regex-like patterns (e.g., `Match("^[0-9]+$", "123abc")` → false)
  - `ContainsWildcard` supports `*` and `?` wildcards (e.g., `ContainsWildcard("hello", "h*l")` → true)

### 2. Numeric Operations
- **Aggregates**: `Range`, `Average`
- **Invariants**:
  - `Range(1, 5)` → [1,2,3,4,5]
  - `Average([]int{1,2,3,4,5})` → 3 (floor division)

###