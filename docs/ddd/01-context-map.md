# Context Map — String Manipulation Bounded Context

## Bounded Contexts
- **String Manipulation** (primary context)
  - Handles UTF-8 string operations, pure functions, and edge cases
  - Includes: Split, EncodeHex, Range, ParseInt, Match, ContainsWildcard, Log

## Aggregate Boundaries
- Each string operation is a standalone aggregate (pure functions with no side effects)
- Surrogate pairs and Unicode handling are context-specific invariants

## Domain Events/Invariants
- **UTF-8 Consistency**: All operations must treat input as valid UTF-8
- **Pure Functions**: No I/O, no clocks, no network (as per specs/dsl.md)
- **Surrogate Pair Handling**: Split must treat surrogate pairs as single code points
- **Inclusive Ranges**: Range(start, end) includes both endpoints
- **Wildcard Matching**: ContainsWildcard supports * and ? patterns

## Open Questions
- How to explicitly validate surrogate pair splitting in Split tests?
-