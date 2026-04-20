# Context Map

## Bounded Contexts
- **String Manipulation** (domain/split.go, domain/reverse.go)
  - Aggregate Roots: `Split`, `Reverse`
  - Invariants:
    - `Split("", "")