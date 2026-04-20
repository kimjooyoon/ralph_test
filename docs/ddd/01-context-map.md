# Context Map

## Bounded Context: String Manipulation
- **Domain Model**: 
  - `Split(s string, sep string) []string` - UTF-8 byte splitting (for empty sep)
  - `Reverse(s string) string` - UTF-8 aware reversal
  - `Join(ss []string, sep string) string` - UTF-8 safe concatenation
  - `Trim(s string) string` - UTF-8 aware trimming
  - `Repeat(s string, count int) string` - UTF-8 safe repetition
  - `Replace(s, old, new string) string` - UTF-8 aware replacement
  - `Echo(s string) string` - identity function

## Aggregate Boundaries
- Each string operation is a stateless function (no aggregate roots)
- UTF-8 handling is a cross-cutting concern across all operations

## Domain Events
- No explicit events (pure functions with no side effects)

## Invariants
- `Split` with empty separator splits by UTF-8 bytes (not code points)
- `Reverse` handles surrogate pairs and combining marks correctly
- All operations must work with valid UTF-8 input
- Functions must be pure (no I/O, no clocks, no network)

## Open Questions
- Should `Split` with empty separator be a separate bounded context?
- How to handle invalid UTF-8 input? (Currently not specified)
