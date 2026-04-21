# Context Map

## Bounded Contexts

### 1. StringHelpers
- **Aggregate Roots**: `Split`, `Join`, `Reverse`, `Trim`, `Replace`, `ContainsWildcard`, `Match`
- **Invariants**:
  - `Split` must split UTF